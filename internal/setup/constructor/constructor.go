package constructor

import (
	"github.com/phn00dev/go-task-manager/internal/app"
	authConstructor "github.com/phn00dev/go-task-manager/internal/domain/auth/constructor"
	userConstructor "github.com/phn00dev/go-task-manager/internal/domain/user/constructor"

)

func Build(dependencies *app.Dependencies) {
	authConstructor.InitAuthRequirements(dependencies.DB)
	userConstructor.InitUserRequirements(dependencies.DB)
}
