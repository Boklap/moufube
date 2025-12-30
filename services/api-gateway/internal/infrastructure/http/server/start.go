package server

import (
	"net/http"
	"os"
)

func StartHTTP(httpServer *http.Server) {
	err := httpServer.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}
