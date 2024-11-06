package redis

import "github.com/go-redis/redis"

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func Connect(cfg *RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port,
	})

	if err := client.Ping().Err(); err != nil {
		panic(err)
	}

	return client
}
