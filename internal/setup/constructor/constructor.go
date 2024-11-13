package constructor

import (
	"github.com/phn00dev/go-task-manager/internal/app"
	adminConstructor "github.com/phn00dev/go-task-manager/internal/domain/admin/constructor"
	teamConstructor "github.com/phn00dev/go-task-manager/internal/domain/team/constructor"
	userConstructor "github.com/phn00dev/go-task-manager/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	// admin
	adminConstructor.InitAdminRequirements(dependencies.DB)
	userConstructor.InitUserRequirements(dependencies.DB)
	teamConstructor.InitTeamRequirements(dependencies.DB)
}
