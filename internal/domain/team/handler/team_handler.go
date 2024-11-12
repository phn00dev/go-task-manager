package handler

import (
	"github.com/gin-gonic/gin"
)

type TeamHandler interface {
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
