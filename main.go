package main

import (
	"log/slog"
	"net/http"
	"os"
	"text/template"
)

func main() {

	// new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// UHC template cache
	tmpl, err := template.ParseFiles("uhc.tmpl")
	if err != nil {
		logger.Error(err.Error())
	}

	api := api{
		logger:        logger,
		templateCache: tmpl,
	}

	// Init server struct
	srv := http.Server{
		Addr:    ":8080",
		Handler: api.routes(),
	}

	// Start main server
	logger.Info("Starting server")

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
