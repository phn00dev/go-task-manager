package service

import (
	"errors"

	"github.com/phn00dev/go-task-manager/internal/domain/admin/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/admin/repository"
	"github.com/phn00dev/go-task-manager/internal/models"
	passwordhash "github.com/phn00dev/go-task-manager/internal/utils/password_hash"
	admintoken "github.com/phn00dev/go-task-manager/pkg/jwt_token/adminToken"
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
	admins, err := adminService.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*models.Admin, error) {
	if adminID == 0 {
		return nil, errors.New("admin id not zero")
	}
	admin, err := adminService.adminRepo.GetOne(adminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (adminService adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) error {
	// password barlag
	if createRequest.Password != createRequest.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}
	// generate password
	passwordHash := passwordhash.GeneratePassword(createRequest.Password)

	createAdmin := models.Admin{
		Firstname:    createRequest.Firstname,
		Lastname:     createRequest.Lastname,
		Email:        createRequest.Email,
		PasswordHash: passwordHash,
		AdminRole:    "admin",
	}
	return adminService.adminRepo.Create(createAdmin)
}

func (adminService adminServiceImp) UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error {
	if adminID == 0 {
		return errors.New("admin id not zero")
	}
	admin, err := adminService.adminRepo.GetOne(adminID)
	if err != nil {
		return err
	}

	admin.Firstname = updateRequest.Firstname
	admin.Lastname = updateRequest.Lastname
	admin.Email = updateRequest.Email

	// update admin data
	return adminService.adminRepo.Update(admin.ID, *admin)
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	if adminID == 0 {
		return errors.New("admin id not zero")
	}
	admin, err := adminService.adminRepo.GetOne(adminID)
	if err != nil {
		return err
	}
	return adminService.adminRepo.Delete(admin.ID)
}

func (adminService adminServiceImp) LoginAdmin(loginRequest dto.LoginRequest) (*dto.LoginAdminResponse, error) {
	// get admin
	admin, err := adminService.adminRepo.GetAdminByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	// check password
	if err := passwordhash.CheckPasswordHash(loginRequest.Password, admin.PasswordHash); err != nil {
		return nil, err
	}

	// generate token
	accessToken, err := admintoken.GenerateAdminToken(admin.ID, admin.AdminRole)
	if err != nil {
		return nil, err
	}
	// login response
	loginResponse := dto.NewLoginAdminResponse(admin, *accessToken)

	return loginResponse, nil
}
