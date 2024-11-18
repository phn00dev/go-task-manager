package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	UpdateData(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	DeleteProfile(ctx *gin.Context)
}
