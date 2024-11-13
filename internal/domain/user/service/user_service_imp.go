package service

import (
	"errors"

	"github.com/phn00dev/go-task-manager/internal/domain/user/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/user/repository"
	"github.com/phn00dev/go-task-manager/internal/models"
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

func (userService userServiceImp) GetAll() ([]models.User, error) {
	users, err := userService.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService userServiceImp) Getone(userID int) (*models.User, error) {
	if userID == 0 {
		return nil, errors.New("user ID must not be zero")
	}
	user, err := userService.userRepo.GetOne(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService userServiceImp) GetProfile(userID int) (*dto.UserResponse, error) {
	if userID == 0 {
		return nil, errors.New("user id not null")
	}
	user, err := userService.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserResponse(user)
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
	user.Firstname = updateRequest.Firstname
	user.Lastname = updateRequest.Lastname
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
	if err := passwordhash.CheckPasswordHash(request.OldPassword, user.PasswordHash); err != nil {
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
