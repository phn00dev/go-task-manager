package dto

import (
	"github.com/phn00dev/go-task-manager/internal/models"

)

type UserProfileResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LastLogin string `json:"last_login"`
}

func NewUserProfileResponse(user *models.User) *UserProfileResponse {
	return &UserProfileResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("02-01-2006 15:04:05"),
		LastLogin: user.LastLogin.Format("02-01-2006 15:04:05"),
	}
}
