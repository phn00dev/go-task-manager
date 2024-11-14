package service

import (
	"errors"

	"github.com/phn00dev/go-task-manager/internal/domain/team/dto"
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
	teams, err := teamService.teamRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (teamService teamServiceImp) GetOneTeam(teamID int) (*models.Team, error) {
	if teamID == 0 {
		return nil, errors.New("team ID must not be zero")
	}
	team, err := teamService.teamRepo.GetOne(teamID)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (teamService teamServiceImp) CreateTeam(createRequest dto.CreateTeamRequest) error {
	createTeam := models.Team{
		TeamName:   createRequest.TeamName,
		TeamStatus: createRequest.TeamStatus,
		AdminID:    createRequest.AdminID,
	}
	return teamService.teamRepo.Create(createTeam)
}

func (teamService teamServiceImp) UpdateTeam(teamID int, updateRequest dto.UpdateTeamRequest) error {
	if teamID == 0 {
		return errors.New("team ID must not be zero")
	}
	team, err := teamService.teamRepo.GetOne(teamID)
	if err != nil {
		return err
	}
	team.TeamName = updateRequest.TeamName
	team.TeamStatus = updateRequest.TeamStatus
	team.AdminID = updateRequest.AdminID
	return teamService.teamRepo.Update(team.ID, *team)

}

func (teamService teamServiceImp) DeleteTeam(teamID int) error {
	if teamID == 0 {
		return errors.New("team ID must not be zero")
	}
	team, err := teamService.teamRepo.GetOne(teamID)
	if err != nil {
		return err
	}
	return teamService.teamRepo.Delete(team.ID)
}
