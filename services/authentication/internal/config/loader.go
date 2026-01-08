package config

import (
	"moufube.com/m/internal/appctx/env"
	"moufube.com/m/internal/appctx/strings"
)

func StringLoader(field *string, key string) fieldLoader {
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

func IntLoader(field *int, key string) fieldLoader {
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

func Int64Loader(field *int64, key string) fieldLoader {
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
