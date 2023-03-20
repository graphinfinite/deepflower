package app

import (
	"context"
	"deepflower/config"
	ctrl "deepflower/internal/controllers"
	"deepflower/internal/repository"
	"deepflower/internal/usecase"
	"deepflower/pkg/postgres"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"github.com/go-chi/cors"
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
	dbPool, err := postgres.NewPostgresPool(cfg.Db.Psql)
	if err != nil {
		return err
	}
	defer dbPool.Close()
	zlog.Info().Msgf("migrate... ")
	//TODO temporary
	if err := postgres.MigrateDb(dbPool); err != nil {
		return err
	}

	client := http.Client{Timeout: time.Second * 10}
	// auth
	userstore := repository.NewUserStorage(dbPool)
	authUC := usecase.NewAuthUsecase(
		&userstore,
		cfg.Auth.Cost,
		cfg.Auth.Signing_key,
		time.Duration(cfg.Auth.Token_ttl)*time.Minute)
	auth := ctrl.NewAuthController(&authUC, &zlog)
	bot, _ := ctrl.NewBot(false, cfg.Telegram.Token, &client, &authUC, &zlog)

	//user
	userUC := usecase.NewUserUC(&userstore)
	user := ctrl.NewUserController(&userUC, &zlog)

	// dream
	dreamstore := repository.NewDreamStorage(dbPool)
	dreamUC := usecase.NewDreamUsecase(&dreamstore)
	dream := ctrl.NewDreamController(&dreamUC, &zlog)
	//task := ctrl.NewTaskController(&zlog)
	//
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello")) })
	r.Post("/bot", bot.TelegramBotMessageReader)
	r.Post("/auth/sign-in", auth.Login)
	r.Route("/user", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Get("/", user.GetUserInfo)
	})
	r.Route("/dreams", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Post("/", dream.CreateDream)
		r.Get("/", dream.GetAllUserDreams)
		r.Patch("/{dreamId}", dream.UpdateUserDream)
		r.Delete("/{dreamId}", dream.DeleteUserDream)
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
