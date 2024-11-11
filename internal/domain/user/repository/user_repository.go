package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CheckUserByEmail(email string) (*models.User, error)
	GetUserByID(userID int) (*models.User, error)
	Update(userID int, user models.User) error
	UpdatePassword(userID int, newPassword string) error
	Delete(userID int) error
}
