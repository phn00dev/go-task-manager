package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetOne(adminID int) (*models.Admin, error)
	Create(admin models.Admin) error
	Update(adminID int, admin models.Admin) error
	Delete(adminID int) error
}
