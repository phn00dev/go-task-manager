package service

import (
	"github.com/phn00dev/go-task-manager/internal/domain/team/repository"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type teamServiceImp struct {
	teamRepo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) TeamService {
	return teamServiceImp{
		teamRepo: repo,
	}
}

func (teamService teamServiceImp) GetAllTeams() ([]models.Team, error) {
	panic("team service imp")
}

func (teamService teamServiceImp) GetOneTeam(teamID int) (*models.Team, error) {
	panic("team service imp")
}

func (teamService teamServiceImp) CreateTeam(team models.Team) error {
	panic("team service imp")
}

func (teamService teamServiceImp) UpdateTeam(teamID int, team models.Team) error {
	panic("team service imp")
}

func (teamService teamServiceImp) DeleteTeam(teamID int) error {
	panic("team service imp")
}
