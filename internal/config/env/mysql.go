package env

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type mysqlEnvConfig struct {
	Host     string `env:"MYSQL_HOST,required"`
	Port     string `env:"MYSQL_PORT,required"`
	ExtPort  string `env:"MYSQL_EXT_PORT,required"`
	User     string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Database string `env:"MYSQL_DATABASE,required"`
}

type mysqlConfig struct {
	cfg mysqlEnvConfig
}

func NewMysqlConfig() (*mysqlConfig, error) {
	var cfg mysqlEnvConfig
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &mysqlConfig{cfg: cfg}, nil
}

func (p *mysqlConfig) URI() string {
	return fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s",
		p.cfg.User,
		p.cfg.Password,
		p.cfg.Database,
	)
}
