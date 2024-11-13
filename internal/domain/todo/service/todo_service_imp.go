package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/todo/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/todo/repository"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type todoServiceImp struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoServiceImp{
		todoRepo: todoRepo,
	}
}

func (t todoServiceImp) GetAllTodos() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoServiceImp) GetOneTodo(todoID int) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoServiceImp) CreateTodo(createRequest dto.CreateTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (t todoServiceImp) UpdateTodo(todoID int, updateRequest dto.UpdateTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (t todoServiceImp) DeleteTodo(todoID int) error {
	//TODO implement me
	panic("implement me")
}
