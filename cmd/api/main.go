package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.NewJSONHandler(os.Stdout, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}

type Scenario struct {
	ID          int      `json:"id"`
	Tittel      string   `json:"title"`
	Beskrivelse string   `json:"description"`
	Temaer      []string `json:"themes"`
}

func home(w http.ResponseWriter, r *http.Request) {
	scenarios := []Scenario{
		{
			ID:          1,
			Tittel:      "Feilkonfigurasjon i IAM gir uautorisert tilgang",
			Beskrivelse: "En ny IAM-policy ble rullet ut uten test. Dette ga tilgang til sensitive data.",
			Temaer:      []string{"IAM", "Personvern"},
		},
		{
			ID:          2,
			Tittel:      "Endring i skytjeneste uten rollback",
			Beskrivelse: "En oppgradering førte til driftsstans uten mulighet for reversering.",
			Temaer:      []string{"Endringshåndtering", "Drift"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scenarios)
}
