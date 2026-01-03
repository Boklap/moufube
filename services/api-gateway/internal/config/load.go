package config

import (
	"moufube.com/m/internal/appctx/env"
	"moufube.com/m/internal/appctx/strings"
	"moufube.com/m/internal/apperr"
)

func Load() (*Config, error) {
	cfg := &Config{}
	loaders := createLoaders(cfg)

	for _, l := range loaders {
		if err := l.load(); err != nil {
			return nil, apperr.FailToLoadConfig(err)
		}
	}

	return cfg, nil
}

func createLoaders(cfg *Config) []fieldLoader {
	strLoader := func(field *string, key string) fieldLoader {
		return fieldLoader{
			load: func() error {
				v, err := env.Get(key)
				if err != nil {
					return err
				}
				*field = v
				return nil
			},
		}
	}

	intLoader := func(field *int, key string) fieldLoader {
		return fieldLoader{
			load: func() error {
				v, err := env.Get(key)
				if err != nil {
					return err
				}
				*field, err = strings.ToInt(v)
				return err
			},
		}
	}

	int64Loader := func(field *int64, key string) fieldLoader {
		return fieldLoader{
			load: func() error {
				v, err := env.Get(key)
				if err != nil {
					return err
				}
				*field, err = strings.ToInt64(v)
				return err
			},
		}
	}

	return []fieldLoader{
		strLoader(&cfg.Environment, "ENVIRONMENT"),
		int64Loader(&cfg.ReadTimeout, "READ_TIMEOUT"),
		int64Loader(&cfg.WriteTimeout, "WRITE_TIMEOUT"),
		int64Loader(&cfg.IdleTimeout, "IDLE_TIMEOUT"),
		int64Loader(&cfg.ShutdownTimeout, "SHUTDOWN_TIMEOUT"),
		int64Loader(&cfg.MaxHeaderBytes, "MAX_HEADER_BYTES"),
		int64Loader(&cfg.MinMultipartMemory, "MIN_MULTIPART_MEMORY"),
		int64Loader(&cfg.MaxMultipartMemory, "MAX_MULTIPART_MEMORY"),
		intLoader(&cfg.HTTPPort, "HTTP_PORT"),
		intLoader(&cfg.SizeIdentityToken, "SIZE_IDENTITY_TOKEN"),
		intLoader(&cfg.VisitorTokenExpireDays, "VISITOR_TOKEN_EXPIRE_DAYS"),
		strLoader(&cfg.RedisHost, "REDIS_HOST"),
		strLoader(&cfg.RedisPort, "REDIS_PORT"),
		strLoader(&cfg.RedisPassword, "REDIS_PASSWORD"),
		intLoader(&cfg.IdentityDB, "IDENTITY_DB"),
	}
}
