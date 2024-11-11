package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/internal/domain/auth/handler"
	"github.com/phn00dev/go-task-manager/internal/domain/auth/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/auth/service"

)

var (
	authRepo    repository.AuthRepository
	authService service.AuthService
	AuthHandler handler.AuthHandler
)

func InitAuthRequirements(db *gorm.DB) {
	authRepo = repository.NewAuthRepository(db)
	authService = service.NewAuthService(authRepo)
	AuthHandler =handler.NewAuthHandler(authService)
}
