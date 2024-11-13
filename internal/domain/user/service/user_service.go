package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/user/dto"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type UserService interface {
	//admin for user services
	GetAll() ([]models.User, error)
	Getone(userID int) (*models.User, error)
	GetProfile(userID int) (*dto.UserResponse, error)
	UpdateUser(userID int, updateRequest dto.UpdateUserRequest) error
	UpdateUserPassword(userID int, request dto.UpdateUserPassword) error
	DeleteUser(userID int) error
}
