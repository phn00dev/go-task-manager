package models

import "time"

type TodoStatus string

const (
	New       TodoStatus = "new"
	Working   TodoStatus = "working"
	Done      TodoStatus = "done"
	Paused    TodoStatus = "paused"
	Cancelled TodoStatus = "cancelled"
	Overdue   TodoStatus = "overdue"
)

// Todo struct represents the todos table
type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"todo_title"`
	Description string     `json:"todo_description"`
	Status      TodoStatus `json:"todo_status"`
	UserID      int        `json:"user_id"`
	TeamID      *int       `json:"team_id,omitempty"` // TeamID can be null
	Deadline    *time.Time `json:"deadline,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
