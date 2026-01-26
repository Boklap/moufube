package logger

import "log/slog"

func InitSlog() *slog.Logger {
	return slog.Default()
}
