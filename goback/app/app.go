package app

import (
	"context"
	"deepflower/config"
	ctrl "deepflower/internal/controllers"
	"deepflower/internal/repository"
	"deepflower/internal/usecase"
	"deepflower/pkg/postgres"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"deepflower/pkg/telegram"

	"deepflower/internal/observer"

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

	zlog.Info().Msgf("start with config: \n%+v", cfg)
	zlog.Info().Msgf("connect to db...")
	dbPool, err := postgres.NewPostgresPool(cfg.Db.Psql)
	if err != nil {
		zlog.Err(err).Msg("App/NewPostgresPool ")
		return err
	}
	defer dbPool.Close()

	// TODO migrations
	zlog.Info().Msgf("up migrations... ")
	if err := postgres.MigrateUp(dbPool); err != nil {
		zlog.Err(err).Msg("App/MigrateUp ")
		return err
	}

	// Create publisher1 - tg bot
	zlog.Info().Msgf("start telegram bot... ")
	client := http.Client{Timeout: time.Second * 10}
	bot, err := telegram.NewBot(cfg.Telegram.Token, cfg.Telegram.Buffer, client, cfg.Telegram.Debug, &zlog)
	if err != nil {
		zlog.Err(err).Msg("App/NewBot ")
		return err
	}

	botOutChan := make(chan observer.Event, cfg.Telegram.Buffer)
	defer bot.Bot.StopReceivingUpdates()
	go func(botOutChan chan observer.Event) {
		bot.StartReceiveUpdates(0, 100, 10, botOutChan)
	}(botOutChan)

	// OBS
	obs := observer.NewObserver(&zlog)
	// OBS
	obs.AddPublisherChain(botOutChan)

	// Auth
	authUC := usecase.NewAuthUsecase(
		repository.NewUserStorage(dbPool),
		cfg.Auth.Cost,
		cfg.Auth.Signing_key,
		time.Duration(cfg.Auth.Token_ttl)*time.Minute)
	auth := ctrl.NewAuthController(authUC, &bot, &zlog)

	// OBS
	obs.AddTopicsHandler([]observer.Topic{}, auth.Registration)

	// User
	user := ctrl.NewUserController(usecase.NewUserUC(repository.NewUserStorage(dbPool)), &zlog)

	// Dream
	dream := ctrl.NewDreamController(usecase.NewDreamUsecase(repository.NewDreamStorage(dbPool)), &zlog)

	// Location
	loc := ctrl.NewLocationController(usecase.NewLocationUsecase(repository.NewLocationStorage(dbPool)), &zlog)

	// Project

	projstore := repository.NewProjectStorage(dbPool)
	projUC := usecase.NewProjectUsecase(projstore)
	proj := ctrl.NewProjectController(projUC, &zlog)

	// Router settings
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
	r.Post("/auth/sign-in", auth.Login)
	r.Route("/user", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Get("/", user.GetUserInfo)
	})
	r.Route("/dreams", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Post("/", dream.CreateDream)
		r.Get("/", dream.SearchDreams)
		r.Patch("/{dreamId}", dream.UpdateUserDream)
		r.Delete("/{dreamId}", dream.DeleteUserDream)
		r.Post("/{dreamId}/publish", dream.PublishDream)
		r.Post("/{dreamId}/energy", dream.AddEnergyToDream)
	})

	r.Route("/locations", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Post("/", loc.CreateLocation)
		r.Get("/", loc.SearchLocations)
		r.Patch("/{locationId}", loc.UpdateUserLocation)
		r.Delete("/{locationId}", loc.DeleteUserLocation)
		r.Post("/{locationId}/energy", loc.AddEnergyToLocation)
		r.Get("/{locationId}/dreams", loc.GetLocationDreams)
	})

	r.Route("/projects", func(r chi.Router) {
		r.Use(auth.JWT)
		r.Post("/", proj.CreateProject)
		r.Get("/", proj.SearchProjects)
		r.Patch("/{projectId}", proj.UpdateUserProject)
		r.Delete("/{projectId}", proj.DeleteUserProject)
		r.Post("/{projectId}/publish", proj.PublishProject)
		//r.Post("/{projectId}/energy", proj.AddEnergyToProject)

		// TODO
		r.Post("/{projectId}/node/{nodeId}/addenergy", proj.AddEnergyToTask)
		r.Post("/{projectId}/node/{nodeId}/grab", proj.ToWorkTask)
		r.Post("/{projectId}/node/{nodeId}/close", proj.CloseTask)
	})

	// HTTP Server
	app.httpServer = &http.Server{
		Addr:           net.JoinHostPort(cfg.Server.Host, cfg.Server.Port),
		Handler:        r,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	zlog.Info().Msgf("deepflower server start... %s", app.httpServer.Addr)
	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			zlog.Err(err).Msg("💀💀💀")
		}
		zlog.Info().Msg("💀")

	}()

	// Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return app.httpServer.Shutdown(ctx)
}
