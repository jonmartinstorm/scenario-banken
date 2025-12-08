package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /scenario/view/{id}", getScenarioView)
	mux.HandleFunc("GET /scenario/create", getScenarioCreate)
	mux.HandleFunc("POST /scenario/create", postScenarioCreate)

	logger.Info("starting server", slog.String("addr", ":4000"))
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
