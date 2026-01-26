package writer

import "github.com/redis/go-redis/v9"

func NewIdentityWriterImpl(rdb *redis.Client) *IdentityWriterImpl {
	return &IdentityWriterImpl{
		rdb: rdb,
	}
}
