package config

import (
	"log"
	"task_manager/internal/config/env"

	"github.com/joho/godotenv"
)

var appConfig *config

type config struct {
	Http   HttpConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
	Logger LoggerConfig
}

func Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	httpCfg, err := env.NewHTTPConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}

	mysqlCfg, err := env.NewMysqlConfig()
	if err != nil {
		log.Println("failed to load order http config")
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Http:   httpCfg,
		MySQL:  mysqlCfg,
		Redis:  redisCfg,
		Logger: loggerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
