package reader

import (
	"github.com/redis/go-redis/v9"
)

func NewIdentityReader(rdb *redis.Client) *IdentityReaderImpl {
	return &IdentityReaderImpl{
		rdb: rdb,
	}
}
