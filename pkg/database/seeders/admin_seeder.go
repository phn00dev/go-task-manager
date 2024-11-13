package seeders

import (
	"fmt"
	"github.com/phn00dev/go-task-manager/internal/models"
	passwordhash "github.com/phn00dev/go-task-manager/internal/utils/password_hash"
)

func (dbSeeder DBSeeder) adminSeeder() error {
	var admins []models.Admin
	passwordHash := passwordhash.GeneratePassword("12345678")
	superAdmin := models.Admin{
		Firstname:    "Hudayberdi",
		Lastname:     "Polatov",
		Email:        "polat@gmail.com",
		PasswordHash: passwordHash,
		AdminRole:    "super_admin",
	}
	admins = append(admins, superAdmin)
	for i := 1; i <= 10; i++ {
		admin := models.Admin{
			Firstname:    fmt.Sprintf("admin-%d", i),
			Lastname:     fmt.Sprintf("admin-%d", i),
			Email:        fmt.Sprintf("admin-%d@gmail.com", i),
			PasswordHash: passwordHash,
			AdminRole:    "admin",
		}
		admins = append(admins, admin)
	}
	if err := dbSeeder.db.Create(&admins).Error; err != nil {
		return err
	}
	return nil
}
