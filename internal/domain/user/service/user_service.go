package service

import "github.com/phn00dev/go-task-manager/internal/domain/user/dto"

type UserService interface {
	GetProfile(userID int) (*dto.UserProfileResponse, error)
	UpdateUser(userID int, updateRequest dto.UpdateUserRequest) error
	UpdateUserPassword(userID int, request dto.UpdateUserPassword) error
	DeleteUser(userID int) error
}
