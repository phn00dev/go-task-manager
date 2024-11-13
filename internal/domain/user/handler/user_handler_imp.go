package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/user/service"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) GetUserProfile(ctx *gin.Context) {

}

func (userHandler userHandlerImp) UpdateData(ctx *gin.Context) {

}

func (userHandler userHandlerImp) UpdatePassword(ctx *gin.Context) {

}

func (userHandler userHandlerImp) DeleteProfile(ctx *gin.Context) {

}
