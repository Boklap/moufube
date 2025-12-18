package server

import (
	"log/slog"
	"net/http"
	"os"
)

func StartHTTP(httpServer *http.Server) {
	err := httpServer.ListenAndServe()
	if err != nil {
		slog.Error("Can not run http server")
		os.Exit(1)
	}

	slog.Info("HTTP server is running.")
}
