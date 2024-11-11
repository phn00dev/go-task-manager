package service

import (
	"errors"

	"github.com/phn00dev/go-task-manager/internal/domain/user/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/user/repository"
	passwordhash "github.com/phn00dev/go-task-manager/internal/utils/password_hash"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) GetProfile(userID int) (*dto.UserProfileResponse, error) {
	if userID == 0 {
		return nil, errors.New("user id not null")
	}
	user, err := userService.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserProfileResponse(user)
	return userResponse, nil
}

func (userService userServiceImp) UpdateUser(userID int, updateRequest dto.UpdateUserRequest) error {
	if userID == 0 {
		return errors.New("user ID must be non-zero")
	}
	user, err := userService.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("user  not found")
	}
	user.Name = updateRequest.Name
	user.Email = updateRequest.Email
	return userService.userRepo.Update(user.ID, *user)
}

func (userService userServiceImp) UpdateUserPassword(userID int, request dto.UpdateUserPassword) error {
	if userID == 0 {
		return errors.New("user id not null")
	}

	if request.Password != request.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	user, err := userService.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	// check password
	if err := passwordhash.CheckPasswordHash(request.OldPassword, user.Password); err != nil {
		return errors.New("old password wrong")
	}

	// hash password generate
	hashPassword := passwordhash.GeneratePassword(request.Password)
	// update password
	return userService.userRepo.UpdatePassword(user.ID, hashPassword)

}

func (userService userServiceImp) DeleteUser(userID int) error {
	if userID == 0 {
		return errors.New("user ID must be non-zero")
	}
	user, err := userService.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	return userService.userRepo.Delete(user.ID)
}
