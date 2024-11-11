package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/models"
)

type authRepositoryImp struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return authRepositoryImp{
		DB: db,
	}
}

func (authRepo authRepositoryImp) Create(user models.User) error {
	panic("auth repo imp")
}

func (authRepo authRepositoryImp) GetByEmail(email string) (*models.User, error) {
	panic("auth repo imp")
}
