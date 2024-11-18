package models

import (
	"github.com/phn00dev/go-task-manager/internal/constants"
	"time"
)

type Todo struct {
	ID          int                  `json:"id"`
	Title       string               `json:"todo_title"`
	Description string               `json:"todo_description"`
	Status      constants.TodoStatus `json:"todo_status"`
	UserID      int                  `json:"user_id"`
	TeamID      *int                 `json:"team_id"`
	Deadline    *time.Time           `json:"deadline"`
	CompletedAt *time.Time           `json:"completed_at"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
