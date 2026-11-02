package main

import (
	"context"
	"os/signal"
	"syscall"
	"task_manager/internal/app"
	"task_manager/internal/config"
	"task_manager/platform/closer"
	"task_manager/platform/pkg/logger"
	"time"

	"go.uber.org/zap"
)

const configPath = "../.env"

func main() {
	if err := config.Load(configPath); err != nil {
		logger.Error(context.Background(), "failed to load iam config: %v", zap.Error(err))
		return
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()
	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.New(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать iam приложение", zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе iam приложения", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}
