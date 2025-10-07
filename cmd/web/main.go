package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /scenario/view/{id}", getScenarioView)
	mux.HandleFunc("GET /scenario/create", getScenarioCreate)
	mux.HandleFunc("POST /scenario/create", postScenarioCreate)

	err := http.ListenAndServe(":4040", mux)
	log.Fatal(err)
}
