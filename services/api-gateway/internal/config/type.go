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
}

type fieldLoader struct {
	load func() error
}
