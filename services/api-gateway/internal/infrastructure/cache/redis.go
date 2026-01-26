package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/config"
)

func InitCacheConnection(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       cfg.IdentityDB,
	})
}
