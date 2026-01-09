package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /scenario/view/{id}", app.getScenarioView)
	mux.HandleFunc("GET /scenario/create", app.getScenarioCreate)
	mux.HandleFunc("POST /scenario/create", app.postScenarioCreate)

	return mux
}
