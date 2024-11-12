package repository

import "github.com/phn00dev/go-task-manager/internal/models"

type TeamRepository interface {
	GetAll() ([]models.Team, error)
	GetOne(teamID int) (*models.Team, error)
	Create(team models.Team) error
	Update(teamID int, team models.Team) error
	Delete(teamID int) error
}
