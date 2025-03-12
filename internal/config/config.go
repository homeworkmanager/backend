package config

import (
	"errors"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	"homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
)

type Config struct {
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Domain string `envconfig:"DOMAIN"`

	AuthTTL time.Duration `envconfig:"AUTH_TTL"`

	EncodingMode string `envconfig:"ENCODING_MODE"`

	FrontendHost string `envconfig:"FRONTEND_HOST"`
	FrontendPort string `envconfig:"FRONTEND_PORT"`

	postgres.PGConfig
	redis.RedisConfig
}

func NewCfg() *Config {
	var cfg Config

	os.Setenv("TZ", "Europe/Moscow")
	time.Local, _ = time.LoadLocation("Europe/Moscow")

	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	if cfg.Host == "" {
		panic(errors.New("cfg is required"))
	}

	return &cfg
}
