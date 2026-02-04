package writer

import (
	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/config"
)

type IdentityWriterImpl struct {
	rdb *redis.Client
	cfg *config.Config
}
