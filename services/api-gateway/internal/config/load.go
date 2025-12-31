package config

import (
	"moufube.com/m/internal/appctx/env"
	"moufube.com/m/internal/appctx/strings"
	"moufube.com/m/internal/apperr"
)

func Load(env *env.Env, strings *strings.Strings) (*Config, error) {
	cfg := &Config{}
	loaders := createLoaders(env, strings, cfg)

	for _, l := range loaders {
		if err := l.load(); err != nil {
			return nil, apperr.FailToLoadConfig(err)
		}
	}

	return cfg, nil
}

func createLoaders(env *env.Env, strings *strings.Strings, cfg *Config) []fieldLoader {
	return []fieldLoader{
		{
			load: func() error {
				v, err := env.Get("ENVIRONMENT")
				if err != nil {
					return err
				}
				cfg.Environment = v
				return nil
			},
		},
		{
			load: func() error {
				v, err := env.Get("READ_TIMEOUT")
				if err != nil {
					return err
				}
				cfg.ReadTimeout, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("WRITE_TIMEOUT")
				if err != nil {
					return err
				}
				cfg.WriteTimeout, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("IDLE_TIMEOUT")
				if err != nil {
					return err
				}
				cfg.IdleTimeout, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("SHUTDOWN_TIMEOUT")
				if err != nil {
					return err
				}
				cfg.ShutdownTimeout, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("MAX_HEADER_BYTES")
				if err != nil {
					return err
				}
				cfg.MaxHeaderBytes, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("MIN_MULTIPART_MEMORY")
				if err != nil {
					return err
				}
				cfg.MinMultipartMemory, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("MAX_MULTIPART_MEMORY")
				if err != nil {
					return err
				}
				cfg.MaxMultipartMemory, err = strings.ToInt64(v)
				return err
			},
		},
		{
			load: func() error {
				v, err := env.Get("HTTP_PORT")
				if err != nil {
					return err
				}
				cfg.HTTPPort, err = strings.ToInt(v)
				return err
			},
		},
	}
}
