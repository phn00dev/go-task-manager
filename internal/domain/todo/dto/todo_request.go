package dto

import (
	"errors"
	"time"
)

type TodoStatus string

const (
	New       TodoStatus = "new"
	Working   TodoStatus = "working"
	Done      TodoStatus = "done"
	Paused    TodoStatus = "paused"
	Cancelled TodoStatus = "cancelled"
)

type CreateTodoRequest struct {
	Title       string     `json:"todo_title" binding:"required,min=3,max=255"`
	Description string     `json:"todo_description" binding:"omitempty,max=500"`
	Status      TodoStatus `json:"todo_status" binding:"required"`
	UserID      int        `json:"user_id" binding:"required,gt=0"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"todo_title" binding:"omitempty,min=3,max=255"`
	Description string     `json:"todo_description" binding:"omitempty,max=500"`
	Status      TodoStatus `json:"todo_status" binding:"omitempty"`
	UserID      int        `json:"user_id" binding:"omitempty,gt=0"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func (s TodoStatus) IsValid() error {
	switch s {
	case New, Working, Done, Paused, Cancelled:
		return nil
	}
	return errors.New("invalid todo status")
}
