package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/admin/dto"
	"github.com/phn00dev/go-task-manager/internal/domain/admin/service"
	bindandvalidate "github.com/phn00dev/go-task-manager/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-task-manager/internal/utils/response"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: service,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *gin.Context) {
	admins, err := adminHandler.adminService.GetAllAdmins()
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "admins not found", err.Error())
		return
	}
	adminResponses := dto.GetAllAdminResponses(admins)
	response.Success(ctx, http.StatusOK, "all admins", adminResponses)
}

func (adminHandler adminHandlerImp) GetOne(ctx *gin.Context) {
	adminIdStr := ctx.Param("adminID")
	adminId, _ := strconv.Atoi(adminIdStr)
	admin, err := adminHandler.adminService.GetOneAdmin(adminId)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "admin not found", err.Error())
		return
	}
	adminResponse := dto.GetOneAdminResponse(admin)
	response.Success(ctx, http.StatusOK, "get admin", adminResponse)
}

func (adminHandler adminHandlerImp) Create(ctx *gin.Context) {
	var createAdminRequest dto.CreateAdminRequest
	if !bindandvalidate.BindAndValidate(ctx, &createAdminRequest) {
		return
	}

	// create admin
	if err := adminHandler.adminService.CreateAdmin(createAdminRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "admin creation error", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "admin created successfully", nil)
}

func (adminHandler adminHandlerImp) Update(ctx *gin.Context) {
	adminIdStr := ctx.Param("adminID")
	adminId, _ := strconv.Atoi(adminIdStr)

	var updateAdminRequest dto.UpdateAdminRequest
	if !bindandvalidate.BindAndValidate(ctx, &updateAdminRequest) {
		return
	}
	// update admin
	if err := adminHandler.adminService.UpdateAdmin(adminId, updateAdminRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "admin update error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "admin updated successfully", nil)
}

func (adminHandler adminHandlerImp) Delete(ctx *gin.Context) {
	adminIdStr := ctx.Param("adminID")
	adminId, _ := strconv.Atoi(adminIdStr)

	// delete admin
	if err := adminHandler.adminService.DeleteAdmin(adminId); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "admin delete error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "admin delete successfully", nil)
}

// auth admin

func (adminHandler adminHandlerImp) LoginAdmin(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	if !bindandvalidate.BindAndValidate(ctx, loginRequest) {
		return
	}
	// admin login service
	loginResponse, err := adminHandler.adminService.LoginAdmin(loginRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "login error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "admin login successully", loginResponse)
}
