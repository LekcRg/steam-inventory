package cache

import (
	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func New(cfg config.Redis) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Cache{
		client: client,
	}
}
