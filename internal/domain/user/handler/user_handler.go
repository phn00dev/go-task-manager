package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUserProfile(ctx *gin.Context)
	UpdateData(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	DeleteProfile(ctx *gin.Context)
}
