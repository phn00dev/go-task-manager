package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	// admin for user handler
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)

	// user for handlers

	GetUserProfile(ctx *gin.Context)
	UpdateData(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	DeleteProfile(ctx *gin.Context)
}
