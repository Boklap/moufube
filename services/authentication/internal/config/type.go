package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBMS       string

	Environment string
}

type fieldLoader struct {
	load func() error
}
