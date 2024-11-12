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
	panic("team repo imp")
}

func (teamRepo teamRepositoryImp) GetOne(teamID int) (*models.Team, error) {
	panic("team repo imp")
}

func (teamRepo teamRepositoryImp) Create(team models.Team) error {
	panic("team repo imp")
}

func (teamRepo teamRepositoryImp) Update(teamID int, team models.Team) error {
	panic("team repo imp")
}

func (teamRepo teamRepositoryImp) Delete(teamID int) error {
	panic("team repo imp")
}
