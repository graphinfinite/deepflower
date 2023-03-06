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

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(cfg config.Configuration) error {
	zlog := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()

	zlog.Info().Msgf("configurate %+v", cfg)
	zlog.Info().Msgf("connect to db...")
	dbPool, err := repository.NewPostgresPool(cfg.Db.Psql)
	if err != nil {
		return err
	}
	defer dbPool.Close()
	zlog.Info().Msgf("migrate... ")
	if err := repository.MigrateDb(dbPool); err != nil {
		return err
	}
	client := http.Client{Timeout: time.Second * 10}
	userstore := repository.NewUserStorage(dbPool)

	authusecase := usecase.NewAuthUsecase(
		&userstore,
		cfg.Auth.Hash_salt,
		cfg.Auth.Signing_key,
		time.Duration(cfg.Auth.Token_ttl)*time.Minute)
	auth := ctrl.NewAuthController(&authusecase, &zlog)
	bot, _ := ctrl.NewBot(false, cfg.Telegram.Token, &client, &authusecase, &zlog)

	// https://api.telegram.org/bot6237215798:AAHQayrhFO8HAvYSi8uVyv4hOcbhJvVr5ro/setWebhook?url=https://dfeb-5-187-87-224.eu.ngrok.io/bot

	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("info")) }) //root
	r.Get("/auth/sign-up/tg", auth.RedirectToTelegram)                                   // redirect to tg bot
	r.Post("/bot", bot.TelegramBotMessageReader)                                         // entrypoint to tg bot
	r.Post("/auth/sign-in", auth.Login)                                                  // jwt-auth
	r.Route("/miau", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("strange")) })
	})

	// HTTP Server
	app.httpServer = &http.Server{
		Addr:           net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	zlog.Info().Msgf("deepflower server start... %s", app.httpServer.Addr)
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
