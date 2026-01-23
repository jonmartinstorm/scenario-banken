package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jonmartinstorm/scenario-banken/internal/models"
)

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	scenarioer, err := app.scenarioer.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Scenarioer = scenarioer

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) getScenarioView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	scenario, err := app.scenarioer.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Scenario = scenario

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) getScenarioCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello fra create"))
}

func (app *application) postScenarioCreate(w http.ResponseWriter, r *http.Request) {
	title := "Superfarlig"
	content := "En superfarlig hendelse \n og masse andre ting \n som newlines"
	scenariotype := "Nasjonal sikkerhet"
	expires := 7

	id, err := app.scenarioer.Insert(title, content, scenariotype, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/scenario/view/%d", id), http.StatusSeeOther)
}
