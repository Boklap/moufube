package config

import (
	"moufube.com/m/internal/shared/types"
	"moufube.com/m/internal/shared/util/env"
)

func Load() *types.Config {
	return &types.Config{
		ReadTimeout:        env.StringToInt64(env.Get("READ_TIMEOUT")),
		WriteTimeout:       env.StringToInt64(env.Get("WRITE_TIMEOUT")),
		IdleTimeout:        env.StringToInt64(env.Get("IDLE_TIMEOUT")),
		MaxHeaderBytes:     env.StringToInt64(env.Get("MAX_HEADER_BYTES")),
		MinMultipartMemory: env.StringToInt64(env.Get("MIN_MULTIPART_MEMORY")),
		MaxMultipartMemory: env.StringToInt64(env.Get("MAX_MULTIPART_MEMORY")),
		HttpPort:           env.StringToInt(env.Get("HTTP_PORT")),
	}
}
