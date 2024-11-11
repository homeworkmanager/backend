package config

import (
	"encoding/json"
	"homeworktodolist/pkg/db/postgres"
	"homeworktodolist/pkg/db/redis"
	"os"
	"time"
)

const (
	cfgPath = "./internal/config/config.json"
)

type Config struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Domain string `json:"domain"`

	AuthTTL time.Duration `json:"auth_ttl"`

	postgres.PGConfig `json:"postgres"`
	redis.RedisConfig `json:"redis"`
}

func NewCfg() *Config {
	jsonFile, err := os.Open(cfgPath)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}

	var config Config

	err = json.NewDecoder(jsonFile).Decode(&config)

	return &config
}
