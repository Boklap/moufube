package bootstrap

import (
	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/infrastructure/repository/identity/reader"
	"moufube.com/m/internal/infrastructure/repository/identity/writer"
	"moufube.com/m/internal/modules/identity/repository"
)

type Repository struct {
	IdentityReader repository.IdentityReader
	IdentityWriter repository.IdentityWriter
}

func InitRepository(rdb *redis.Client) *Repository {
	identityReader := reader.NewIdentityReader(rdb)
	identityWriter := writer.NewIdentityWriterImpl(rdb)

	return &Repository{
		IdentityReader: identityReader,
		IdentityWriter: identityWriter,
	}
}
