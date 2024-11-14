package models

import "time"

type Team struct {
	ID         int       `json:"id"`
	TeamName   string    `json:"team_name"`
	TeamStatus string    `json:"team_status"`
	AdminID    int       `json:"admin_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Admin      Admin     `json:"admin"`
}
