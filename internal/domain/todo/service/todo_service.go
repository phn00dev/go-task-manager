package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/todo/dto"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type TodoService interface {
	GetAllTodos() ([]models.Todo, error)
	GetOneTodo(todoID int) (*models.Todo, error)
	CreateTodo(createRequest dto.CreateTodoRequest) error
	UpdateTodo(todoID int, updateRequest dto.UpdateTodoRequest) error
	DeleteTodo(todoID int) error
}
