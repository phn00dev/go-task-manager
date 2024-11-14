package constructor

import (
	"gorm.io/gorm"

	adminRepository "github.com/phn00dev/go-task-manager/internal/domain/admin/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/team/handler"
	"github.com/phn00dev/go-task-manager/internal/domain/team/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/team/service"
)

var (
	teamRepo    repository.TeamRepository
	adminRepo   adminRepository.AdminRepository
	teamService service.TeamService
	TeamHandler handler.TeamHandler
)

func InitTeamRequirements(db *gorm.DB) {
	teamRepo = repository.NewTeamRepository(db)
	adminRepo = adminRepository.NewAdminRepository(db)
	teamService = service.NewTeamService(teamRepo, adminRepo)
	TeamHandler = handler.NewTeamHandler(teamService)
}
