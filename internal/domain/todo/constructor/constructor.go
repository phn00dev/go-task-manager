package constructor

import (
	"github.com/phn00dev/go-task-manager/internal/domain/todo/handler"
	"github.com/phn00dev/go-task-manager/internal/domain/todo/repository"
	"github.com/phn00dev/go-task-manager/internal/domain/todo/service"
	"gorm.io/gorm"
)

var (
	todoRepo    repository.TodoRepository
	todoService service.TodoService
	TodoHandler handler.TodoHandler
)

func InitTodoRequirements(db *gorm.DB) {
	todoRepo = repository.NewTodoRepository(db)
	todoService = service.NewTodoService(todoRepo)
	TodoHandler = handler.NewTodoHandler(todoService)
}
