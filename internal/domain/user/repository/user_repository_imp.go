package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/models"
)

type userRepositoryImp struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		DB: db,
	}
}

func (userRepo userRepositoryImp) GetUserByID(userID int) (*models.User, error) {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) GetUserByEmail(email string) (*models.User, error) {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) Update(userID int, user models.User) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) UpdatePassword(userID int, newPassword string) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) Delete(userID int) error {
	panic("user repo imp")
}
