package repository

import (
	"errors"

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

func (userRepo userRepositoryImp) GetAll() ([]models.User, error) {
	var users []models.User
	if err := userRepo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepo userRepositoryImp) GetOne(userID int) (*models.User, error) {
	var user models.User
	if err := userRepo.DB.First(&user, userID).Error; err!=nil{
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := userRepo.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) CheckUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := userRepo.DB.Select("users.email").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := userRepo.DB.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) Update(userID int, user models.User) error {
	return userRepo.DB.Model(&models.User{}).Where("id=?", userID).Updates(&user).Error
}

func (userRepo userRepositoryImp) UpdatePassword(userID int, newPassword string) error {
	return userRepo.DB.Model(&models.User{}).Where("id=?", userID).Update("password", newPassword).Error
}

func (userRepo userRepositoryImp) Delete(userID int) error {
	return userRepo.DB.Where("id=?", userID).Delete(&models.User{}).Error
}
