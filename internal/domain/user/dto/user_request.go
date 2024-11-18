package dto

import "github.com/phn00dev/go-task-manager/internal/constants"

type UpdateUserRequest struct {
	Firstname  string           `json:"firstname" binding:"required,min=3,max=50"`
	Lastname   string           `json:"lastname" binding:"required,min=3,max=50"`
	Email      string           `json:"email" binding:"required,email"`
	UserStatus constants.Status `json:"user_status" binding:"required"`
}

type UpdateUserPassword struct {
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
