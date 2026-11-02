package app

import (
	"context"
	"database/sql"
	"task_manager/platform/cache"
	"task_manager/platform/cache/redis"
	"task_manager/platform/closer"

	_ "github.com/go-sql-driver/mysql"

	redigo "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"

	tmV1api "task_manager/internal/api"
	"task_manager/internal/config"
	"task_manager/platform/pkg/logger"
	tmV1 "task_manager/shared/pkg/openapi/tm/v1"
)

type diContainer struct {
	mySQL   *sql.DB
	tmV1Api tmV1.Invoker

	redisPool   *redigo.Pool
	redisClient cache.RedisClient
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) TmV1API(ctx context.Context) tmV1.Invoker {
	if d.tmV1Api == nil {
		d.tmV1Api = tmV1api.NewTaskManagerApi(d.MySQL(ctx))
	}

	return d.tmV1Api
}

func (d *diContainer) MySQL(ctx context.Context) *sql.DB {
	if d.mySQL == nil {
		// Создаем соединение с базой данных
		db, err := sql.Open("mysql", config.AppConfig().MySQL.URI())

		if err != nil {
			logger.Error(ctx, "failed to connect MySQL: %v\n", zap.Error(err))
			return nil
		}
		closer.AddNamed("MySQL client", func(ctx context.Context) error {
			return db.Close()
		})
		// Проверяем, что соединение с базой установлено
		err = db.Ping()
		if err != nil {
			logger.Error(ctx, "failed to ping MySQL: %v\n", zap.Error(err))
			return nil
		}
		logger.Info(ctx, "mysql connect success")
		d.mySQL = db
	}

	return d.mySQL
}

func (d *diContainer) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}

	return d.redisPool
}

func (d *diContainer) RedisClient() cache.RedisClient {
	if d.redisClient == nil {
		d.redisClient = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}

	return d.redisClient
}
