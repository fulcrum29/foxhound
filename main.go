package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	api := api{
		logger: logger,
	}

	// Init server struct
	srv := http.Server{
		Addr:    ":8080",
		Handler: api.routes(),
	}

	// Start main server
	logger.Info("Starting server")

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
