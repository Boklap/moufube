package logger

func Init() *Logger {
	return &Logger{
		Slog: InitSlog(),
	}
}
