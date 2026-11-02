package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"task_manager/internal/config"
	"task_manager/platform/closer"
	"task_manager/platform/pkg/logger"
	"time"

	apiv1 "task_manager/shared/pkg/openapi/tm/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type App struct {
	diContainer *diContainer
	tmAPI       *apiv1.Server
	httpServer  *http.Server
	router      *chi.Mux
}

func New(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.initRouter,
		a.initHTTPServer,
	}
	//Циклично инициализируем зависимости
	for _, dep := range deps {
		if err := dep(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initRouter(ctx context.Context) error {
	r := chi.NewRouter()

	// Добавляем middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	// Монтируем обработчики OpenAPI
	r.Mount("/", a.tmAPI)
	a.router = r

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	tmApi, err := apiv1.NewServer(a.diContainer.TmV1API(ctx), nil)
	if err != nil {
		logger.Error(ctx, "ошибка создания сервера OpenAPI: %v", zap.Error(err))
	}
	a.tmAPI = tmApi

	err = a.initRouter(ctx)
	if err != nil {
		logger.Error(ctx, "ошибка создания router: %v", zap.Error(err))
		return err
	}

	a.httpServer = &http.Server{
		Addr:              config.AppConfig().Http.Address(),
		Handler:           a.router,
		ReadHeaderTimeout: config.AppConfig().Http.ReadTimeOut(),
	}

	closer.AddNamed("HTTP server", func(ctx context.Context) error {
		err := a.httpServer.Shutdown(ctx)
		logger.Info(ctx, ("🛑 Завершение работы сервера..."))
		if err != nil {
			logger.Error(ctx, "❌ Ошибка при остановке сервера: %v\n", zap.Error(err))
			return err
		}
		return nil
	})

	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("🚀 HTTP-сервер запущен на порту %s", config.AppConfig().Http.Port()))

	err := a.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(ctx, "❌ Ошибка запуска сервера:", zap.Error(err))
		return err
	}
	return nil
}

func (a *App) Run(ctx context.Context) error {
	return a.runHTTPServer(ctx)
}
