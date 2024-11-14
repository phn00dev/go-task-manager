package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/models"
)

type teamRepositoryImp struct {
	DB *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return teamRepositoryImp{
		DB: db,
	}
}

func (teamRepo teamRepositoryImp) GetAll() ([]models.Team, error) {
	var teams []models.Team
	if err := teamRepo.DB.Order("id desc").Preload("Admin", func(db *gorm.DB) *gorm.DB {
		return db.Select("id")
	}).Find(&teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}

func (teamRepo teamRepositoryImp) GetOne(teamID int) (*models.Team, error) {
	var team models.Team
	if err := teamRepo.DB.Preload("Admin").First(&team, teamID).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (teamRepo teamRepositoryImp) Create(team models.Team) error {
	return teamRepo.DB.Create(&team).Error
}

func (teamRepo teamRepositoryImp) Update(teamID int, team models.Team) error {
	return teamRepo.DB.Model(&models.Team{}).Where("id=?", teamID).Updates(&team).Error
}

func (teamRepo teamRepositoryImp) Delete(teamID int) error {
	return teamRepo.DB.Delete(&models.Team{}, teamID).Error
}
