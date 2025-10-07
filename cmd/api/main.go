package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello fra scenariobanken"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
