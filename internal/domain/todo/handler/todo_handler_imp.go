package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-task-manager/internal/domain/todo/service"
)

type todoHandlerImp struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) TodoHandler {
	return todoHandlerImp{
		todoService: todoService,
	}
}

func (t todoHandlerImp) GetAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t todoHandlerImp) GetOne(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t todoHandlerImp) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t todoHandlerImp) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t todoHandlerImp) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
