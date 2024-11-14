package service

import (
	"errors"

	adminRepository "github.com/phn00dev/go-task-manager/internal/domain/admin/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/team/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/team/repository"
	"github.com/phn00dev/go-task-manager/internal/models"
)

type teamServiceImp struct {
	teamRepo  repository.TeamRepository
	adminRepo adminRepository.AdminRepository
}

func NewTeamService(repo repository.TeamRepository, adminRepo adminRepository.AdminRepository) TeamService {
	return teamServiceImp{
		teamRepo:  repo,
		adminRepo: adminRepo,
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

	// get admin
	admin, err := teamService.adminRepo.GetOne(createRequest.AdminID)
	if err != nil {
		return errors.New("admin not found")
	}

	createTeam := models.Team{
		TeamName:   createRequest.TeamName,
		TeamStatus: createRequest.TeamStatus,
		AdminID:    admin.ID,
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
	// get admin
	admin, err := teamService.adminRepo.GetOne(updateRequest.AdminID)
	if err != nil {
		return errors.New("admin not found")
	}

	team.TeamName = updateRequest.TeamName
	team.TeamStatus = updateRequest.TeamStatus
	team.AdminID = admin.ID
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
