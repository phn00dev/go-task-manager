package models

import (
	"github.com/phn00dev/go-task-manager/internal/constants"
	"time"
)

type Team struct {
	ID         int              `json:"id"`
	TeamName   string           `json:"team_name"`
	TeamStatus constants.Status `json:"team_status"`
	AdminID    int              `json:"admin_id"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	Admin      Admin            `json:"admin"`
}
