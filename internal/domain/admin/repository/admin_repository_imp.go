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
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) GetOne(adminID int) (*models.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Update(adminID int, admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	panic("admin repo imp")
}
