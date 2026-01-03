package writer

import "github.com/redis/go-redis/v9"

type IdentityWriterImpl struct {
	rdb *redis.Client
}
