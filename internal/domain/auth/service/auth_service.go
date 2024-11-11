package service

import "github.com/phn00dev/go-task-manager/internal/domain/auth/dto"

type AuthService interface {
	RegisterUser(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error)
	LoginUser(loginRequest dto.LoginRequest) (*dto.AuthResponse, error)
}
