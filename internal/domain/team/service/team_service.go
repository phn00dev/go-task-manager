package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/team/dto"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type TeamService interface {
	GetAllTeams() ([]models.Team, error)
	GetOneTeam(teamID int) (*models.Team, error)
	CreateTeam(createRequest dto.CreateTeamRequest) error
	UpdateTeam(teamID int, updateRequest dto.UpdateTeamRequest) error
	DeleteTeam(teamID int) error
}
