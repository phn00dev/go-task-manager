package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/models"
)

type adminRepositoryImp struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		DB: db,
	}
}

func (adminRepo adminRepositoryImp) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	if err := adminRepo.DB.Where("admin_role=?", "admin").
		Select("id", "firstname", "lastname", "admin_role", "email", "created_at", "updated_at", "last_login").
		Order("id desc").
		Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminRepo adminRepositoryImp) GetOne(adminID int) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.DB.Where("admin_role = ?", "admin").First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	return adminRepo.DB.Create(&admin).Error
}

func (adminRepo adminRepositoryImp) Update(adminID int, admin models.Admin) error {
	return adminRepo.DB.Model(&models.Admin{}).Where("id=?", adminID).Updates(admin).Error
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	return adminRepo.DB.Where("id=?", adminID).Delete(&models.Admin{}).Error
}

func (adminRepo adminRepositoryImp) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.DB.Where("email=?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
