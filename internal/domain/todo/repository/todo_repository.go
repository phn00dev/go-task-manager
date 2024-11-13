package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetOne(id int) (*models.Todo, error)
	Create(todo models.Todo) error
	Update(todoID int, todo models.Todo) error
	Delete(todoID int) error
}
