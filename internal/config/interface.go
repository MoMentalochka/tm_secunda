package config

import "time"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}
type HttpConfig interface {
	Address() string
	Port() string
	ReadTimeOut() time.Duration
}
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	CacheTTL() time.Duration
}

type MySQLConfig interface {
	URI() string
}
