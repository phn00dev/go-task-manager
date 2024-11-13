package repository

import (
	"github.com/phn00dev/go-task-manager/internal/models"
	"gorm.io/gorm"
)

type todoRepositoryImp struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return todoRepositoryImp{
		DB: db,
	}
}

func (t todoRepositoryImp) GetAll() ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepositoryImp) GetOne(id int) (*models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepositoryImp) Create(todo models.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (t todoRepositoryImp) Update(todoID int, todo models.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (t todoRepositoryImp) Delete(todoID int) error {
	//TODO implement me
	panic("implement me")
}
