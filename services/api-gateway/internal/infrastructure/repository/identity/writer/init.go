package writer

import (
	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/config"
)

func NewIdentityWriterImpl(
	rdb *redis.Client,
	cfg *config.Config,
) *IdentityWriterImpl {
	return &IdentityWriterImpl{
		rdb: rdb,
		cfg: cfg,
	}
}
