package app

import (
	"context"
	"deepflower/config"
	ctrl "deepflower/internal/controllers"
	"deepflower/internal/repository"
	"deepflower/internal/usecase"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(cfg config.Configuration) error {

	// https://github.com/Permify/go-role
	// https://habr.com/ru/company/vk/blog/692062/

	zlog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Dbname)

	zlog.Printf("psqlInfo: %s", psqlInfo)
	dbPool, err := repository.NewPostgresPool(psqlInfo)
	if err != nil {
		return err
	}
	defer dbPool.Close()
	// migrations
	if err := repository.MigrateDb(dbPool); err != nil {
		return err
	}

	client := http.Client{}

	userstore := repository.NewUserStorage(dbPool)
	authusecase := usecase.NewAuthUsecase(&userstore)
	auth := ctrl.NewAuthController(&authusecase, &zlog)
	bot, _ := ctrl.NewBot(true, &client, &zlog, &authusecase)

	// temporary webhook set
	//_, err := client.Get("https://api.telegram.org/bot6237215798:AAHQayrhFO8HAvYSi8uVyv4hOcbhJvVr5ro/setWebhook?url=https://e33b-109-106-142-77.eu.ngrok.io/bot")
	//if err!= nil{}
	// https://api.telegram.org/bot6237215798:AAHQayrhFO8HAvYSi8uVyv4hOcbhJvVr5ro/setWebhook?url=https://62fb-5-187-75-135.eu.ngrok.io/bot

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("info")) })
	// first auth with telegram account
	r.Get("/auth/register/tg", auth.RedirectToTelegram)
	r.Post("/bot", bot.TelegramBotMessageReader)

	// r.HandleFunc("/auth/login", TelegramBotMessageReader).Methods("POST")
	// r.HandleFunc("/auth/logout", TelegramBotMessageReader).Methods("POST")
	// r.HandleFunc("/auth/remove", TelegramBotMessageReader).Methods("POST")

	// HTTP Server
	app.httpServer = &http.Server{
		Addr:           net.JoinHostPort(viper.GetString("host"), viper.GetString("port")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("deepflower server start... %s", app.httpServer.Addr)
	go func() {
		err := app.httpServer.ListenAndServe()
		zlog.Info().Msg(err.Error())
		fmt.Print("💀")
	}()

	// Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return app.httpServer.Shutdown(ctx)
}
