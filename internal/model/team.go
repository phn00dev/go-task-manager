package model

import "time"

type Team struct {
	ID        int       `json:"id"`
	TeamName  string    `json:"team_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
