package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBMS       string
	DBTimeout  int

	GRPCHost              string
	GRPCTransportProtocol string
	GRPCPort              int

	Environment string
}

type FieldLoader struct {
	load func() error
}
