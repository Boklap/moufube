package types

type Config struct {
	ReadTimeout        int64
	WriteTimeout       int64
	IdleTimeout        int64
	MaxHeaderBytes     int64
	MinMultipartMemory int64
	MaxMultipartMemory int64
	HttpPort           int
}
