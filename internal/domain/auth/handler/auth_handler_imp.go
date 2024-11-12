package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/auth/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/auth/service"
	bindandvalidate "github.com/phn00dev/go-task-manager/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-task-manager/internal/utils/response"
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
	var registerRequest dto.RegisterRequest
	if !bindandvalidate.BindAndValidate(ctx, &registerRequest) {
		return
	}
	// register user
	authResponse, err := authHandler.authService.RegisterUser(registerRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "user register error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "user registered successfully", authResponse)
}

func (authHandler authHandlerImp) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if !bindandvalidate.BindAndValidate(ctx, &loginRequest) {
		return
	}
	// register user
	authResponse, err := authHandler.authService.LoginUser(loginRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "user login error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "user login successfully", authResponse)
}
