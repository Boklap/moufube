package reader

import (
	"github.com/redis/go-redis/v9"
)

type IdentityReaderImpl struct {
	rdb *redis.Client
}
