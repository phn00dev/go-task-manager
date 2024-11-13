package routes

import (
	"github.com/gin-gonic/gin"

	adminConstructor "github.com/phn00dev/go-task-manager/internal/domain/admin/constructor"
)

func AdminRoutes(route *gin.Engine) {
	v1a := route.Group("/v1a/api/admin")
	{
		authAdminRoute := v1a.Group("/auth")
		{
			authAdminRoute.POST("/login", adminConstructor.AdminHandler.LoginAdmin)
		}
		adminRoute := v1a.Group("/admins")
		{
			adminRoute.GET("/", adminConstructor.AdminHandler.GetAll)
			adminRoute.GET("/:adminID", adminConstructor.AdminHandler.GetOne)
			adminRoute.POST("/create", adminConstructor.AdminHandler.Create)
			adminRoute.PUT("/:adminID", adminConstructor.AdminHandler.Update)
			adminRoute.DELETE("/:adminID", adminConstructor.AdminHandler.Delete)
		}
	}
}
