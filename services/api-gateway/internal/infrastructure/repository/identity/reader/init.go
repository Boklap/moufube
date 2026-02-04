package reader

import (
	"github.com/redis/go-redis/v9"
)

func NewIdentityReaderImpl(rdb *redis.Client) *IdentityReaderImpl {
	return &IdentityReaderImpl{
		rdb: rdb,
	}
}
