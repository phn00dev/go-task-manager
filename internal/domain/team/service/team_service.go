package service

import "github.com/phn00dev/go-task-manager/internal/models"

type TeamService interface {
	GetAllTeams() ([]models.Team, error)
	GetOneTeam(teamID int) (*models.Team, error)
	CreateTeam(team models.Team) error
	UpdateTeam(teamID int, team models.Team) error
	DeleteTeam(teamID int) error
}
