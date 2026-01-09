package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /scenario/view/{id}", app.getScenarioView)
	mux.HandleFunc("GET /scenario/create", app.getScenarioCreate)
	mux.HandleFunc("POST /scenario/create", app.postScenarioCreate)

	logger.Info("starting server", slog.String("addr", ":4000"))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
