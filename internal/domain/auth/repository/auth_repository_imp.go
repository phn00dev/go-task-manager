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
	return authRepo.DB.Create(&user).Error
}

func (authRepo authRepositoryImp) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := authRepo.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
