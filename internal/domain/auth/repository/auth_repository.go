package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type AuthRepository interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
}
