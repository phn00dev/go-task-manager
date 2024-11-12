package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/domain/team/handler"
	"github.com/phn00dev/go-task-manager/internal/domain/team/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/team/service"
)

var (
	teamRepo    repository.TeamRepository
	teamService service.TeamService
	TeamHandler handler.TeamHandler
)

func InitTeamRequirements(db *gorm.DB) {
	teamRepo = repository.NewTeamRepository(db)
	teamService = service.NewTeamService(teamRepo)
	TeamHandler = handler.NewTeamHandler(teamService)
}
