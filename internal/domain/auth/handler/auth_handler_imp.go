package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/auth/service"

)

type authHandlerImp struct {
	authService service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return authHandlerImp{
		authService: service,
	}
}

func (authHandler authHandlerImp) Register(ctx *gin.Context) {

}

func (authHandler authHandlerImp) Login(ctx *gin.Context) {

}
