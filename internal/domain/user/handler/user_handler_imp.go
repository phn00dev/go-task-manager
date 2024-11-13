package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/user/service"
	"github.com/phn00dev/go-task-manager/internal/utils/response"

)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) GetAll(ctx *gin.Context) {
	users, err := userHandler.userService.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "users not found", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "all users", users)
}

func (userHandler userHandlerImp) GetOne(ctx *gin.Context) {

}

func (userHandler userHandlerImp) GetUserProfile(ctx *gin.Context) {

}

func (userHandler userHandlerImp) UpdateData(ctx *gin.Context) {

}

func (userHandler userHandlerImp) UpdatePassword(ctx *gin.Context) {

}

func (userHandler userHandlerImp) DeleteProfile(ctx *gin.Context) {

}
