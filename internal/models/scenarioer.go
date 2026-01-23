package models

import (
	"database/sql"
	"errors"
	"time"
)

type Scenario struct {
	ID           int
	Title        string
	Content      string
	ScenarioType string
	Created      time.Time
	Expires      time.Time
}

type ScenarioModel struct {
	DB *sql.DB
}

func (m *ScenarioModel) Insert(title string, content string, scenariotype string, expiresDays int) (int, error) {

	stmt := `INSERT INTO scenarioer (title, content, scenario_type,created, expires)
	VALUES ($1, $2, $3, NOW(), $4)
	RETURNING id`

	expires := time.Now().UTC().Add(time.Duration(expiresDays) * 24 * time.Hour)

	var id int
	err := m.DB.QueryRow(stmt, title, content, scenariotype, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *ScenarioModel) Get(id int) (Scenario, error) {
	stmt := `SELECT id, title, content, scenario_type, created, expires FROM scenarioer
	WHERE expires > NOW() AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	var s Scenario

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.ScenarioType, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Scenario{}, ErrNoRecord
		} else {
			return Scenario{}, err
		}
	}

	return s, nil
}

func (m *ScenarioModel) Latest() ([]Scenario, error) {
	stmt := `SELECT id, title, content, scenario_type, created, expires FROM scenarioer
	WHERE expires > NOW() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var scenarioer []Scenario

	for rows.Next() {
		var s Scenario

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.ScenarioType, &s.Created, &s.Expires)
		if err != nil {
			return nil, err

		}
		scenarioer = append(scenarioer, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return scenarioer, nil
}
