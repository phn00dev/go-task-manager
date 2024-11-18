package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CheckUserByEmail(email string) (*models.User, error)
	Update(userID int, user models.User) error
	Delete(userID int) error
}
