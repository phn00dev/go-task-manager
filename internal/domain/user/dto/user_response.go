package dto

import (
	"github.com/phn00dev/go-task-manager/internal/models"
)

type UserResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LastLogin string `json:"last_login"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("02-01-2006 15:04:05"),
		LastLogin: user.LastLogin.Format("02-01-2006 15:04:05"),
	}
}

func GetAllUserReponses(users []models.User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponse := NewUserResponse(&user)
		userResponses = append(userResponses, *userResponse)
	}
	return userResponses
}
