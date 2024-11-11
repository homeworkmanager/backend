package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func Connect(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}
