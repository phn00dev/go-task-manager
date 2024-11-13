package seeders

import (
	"fmt"
	"github.com/phn00dev/go-task-manager/internal/models"
)

func (dbSeeder DBSeeder) teamSeeder() error {
	var teams []models.Team
	var admins []models.Admin
	// get all admins
	if err := dbSeeder.db.Where("admin_role=?", "admin").Find(&admins).Error; err != nil {
		return err
	}
	for _, admin := range admins {
		team := models.Team{
			TeamName:   fmt.Sprintf("team-%d", admin.ID),
			TeamStatus: "active",
			AdminID:    admin.ID,
		}
		teams = append(teams, team)
	}
	if err := dbSeeder.db.Create(&teams).Error; err != nil {
		return err
	}
	return nil
}
