package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/user/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/user/repository"

)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) GetProfile(userID int) (dto.UserProfileResponse, error) {
	panic("user service imp")
}

func (userService userServiceImp) UpdateUser(userID int, updateRequest dto.UpdateUserRequest) error {
	panic("user service imp")
}

func (userService userServiceImp) UpdateUserPassword(userID int, request dto.UpdateUserPassword) error {
	panic("user service imp")
}

func (userService userServiceImp) DeleteUser(userID int) error {
	panic("user service imp")
}
