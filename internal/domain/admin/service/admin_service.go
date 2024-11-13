package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/admin/dto"
	"github.com/phn00dev/go-task-manager/internal/models"

)

type AdminService interface {
	GetAllAdmins() ([]models.Admin, error)
	GetOneAdmin(adminID int) (*models.Admin, error)
	CreateAdmin(createRequest dto.CreateAdminRequest) error
	UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error
	DeleteAdmin(adminID int) error
	// admin auth
	LoginAdmin(loginRequest dto.LoginRequest) (*dto.LoginAdminResponse, error)
}
