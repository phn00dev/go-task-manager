package constructor

import (
	"github.com/phn00dev/go-task-manager/internal/app"
	teamConstructor "github.com/phn00dev/go-task-manager/internal/domain/team/constructor"
	userConstructor "github.com/phn00dev/go-task-manager/internal/domain/user/constructor"

)

func Build(dependencies *app.Dependencies) {
	userConstructor.InitUserRequirements(dependencies.DB)
	teamConstructor.InitTeamRequirements(dependencies.DB)
}
