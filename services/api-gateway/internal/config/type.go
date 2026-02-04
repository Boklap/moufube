package config

type Config struct {
	Environment            string
	ReadTimeout            int64
	WriteTimeout           int64
	IdleTimeout            int64
	ShutdownTimeout        int64
	MaxHeaderBytes         int64
	MinMultipartMemory     int64
	MaxMultipartMemory     int64
	HTTPPort               int
	SizeIdentityToken      int
	VisitorTokenExpireDays int
	RedisHost              string
	RedisPort              string
	RedisPassword          string
	IdentityDB             int
	RLDuration             int
	RLVisitorMax           int
	RLVisitorMin           int
	GRPCAuthenticationHost string
	GRPCAuthenticationPort int
}

type fieldLoader struct {
	load func() error
}
