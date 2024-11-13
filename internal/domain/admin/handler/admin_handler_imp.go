package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/domain/admin/service"

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
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) GetOne(ctx *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Create(ctx *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Update(ctx *gin.Context) {
	panic("admin handler imp")
}

func (adminHandler adminHandlerImp) Delete(ctx *gin.Context) {
	panic("admin handler imp")
}
