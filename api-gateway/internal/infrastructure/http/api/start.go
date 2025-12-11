package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func Start(httpServer *http.Server) {
	err := httpServer.ListenAndServe()
	if err != nil {
		slog.Error("Can not run http server")
		os.Exit(1)
	}

	fmt.Println("HTTP Server is runnning")
}
