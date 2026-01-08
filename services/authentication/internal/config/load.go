package config

func Load() (*Config, error) {
	cfg := &Config{}
	loaders := createLoaders(cfg)

	for _, l := range loaders {
		if err := l.load(); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func createLoaders(cfg *Config) []fieldLoader {
	return []fieldLoader{
		StringLoader(&cfg.DBHost, "DB_HOST"),
		StringLoader(&cfg.DBPort, "DB_PORT"),
		StringLoader(&cfg.DBUser, "DB_USER"),
		StringLoader(&cfg.DBPassword, "DB_PASSWORD"),
		StringLoader(&cfg.DBName, "DB_NAME"),
		StringLoader(&cfg.DBSSLMode, "DB_SSL_MODE"),
		StringLoader(&cfg.DBMS, "DBMS"),
		StringLoader(&cfg.Environment, "ENVIRONMENT"),
	}
}
