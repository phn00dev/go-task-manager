package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/admin/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/admin/repository"
	"github.com/phn00dev/go-task-manager/internal/models"

)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (adminService adminServiceImp) GetAllAdmins() ([]models.Admin, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*models.Admin, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	panic("admin service imp")
}
