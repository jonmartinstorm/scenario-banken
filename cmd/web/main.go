package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/jonmartinstorm/scenario-banken/internal/models"
	_ "github.com/lib/pq"
)

type application struct {
	logger        *slog.Logger
	scenarioer    *models.ScenarioModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network address")

	dsn := flag.String(
		"dsn",
		"host=localhost port=5432 user=web password=pass dbname=scenariobank sslmode=disable",
		"PostgreSQL data source name",
	)

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger:        logger,
		scenarioer:    &models.ScenarioModel{DB: db},
		templateCache: templateCache,
	}

	logger.Info("starting server", slog.String("addr", ":4000"))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
