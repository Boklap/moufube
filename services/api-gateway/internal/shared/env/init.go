package env

import "log/slog"

func Init(slog *slog.Logger) *Env {
	return &Env{
		slog: slog,
	}
}
