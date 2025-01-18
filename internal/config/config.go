package config

import (
	"errors"
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

	postgres.PGConfig
	redis.RedisConfig
}

func NewCfg() *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	if cfg.Host == "" {
		panic(errors.New("cfg is required"))
	}

	return &cfg
}
