package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/auth/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/auth/repository"
)

type authServiceImp struct {
	authRepo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return authServiceImp{
		authRepo: repo,
	}
}

func (authService authServiceImp) RegisterUser(registerRequest dto.RegisterRequest) (*dto.AuthResponse, error) {
	panic("auth service imp")
}

func (authService authServiceImp) LoginUser(loginRequest dto.LoginRequest) (*dto.AuthResponse, error) {
	panic("auth service imp")
}
