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
	var todos []models.Todo
	if err := t.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (t todoRepositoryImp) GetOne(id int) (*models.Todo, error) {
	var todo models.Todo
	if err := t.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t todoRepositoryImp) Create(todo models.Todo) error {
	return t.DB.Create(&todo).Error
}

func (t todoRepositoryImp) Update(todoID int, todo models.Todo) error {
	return t.DB.Model(&models.Todo{}).Where("id = ?", todoID).Updates(&todo).Error
}

func (t todoRepositoryImp) Delete(todoID int) error {
	return t.DB.Delete(&models.Todo{}, todoID).Error
}
