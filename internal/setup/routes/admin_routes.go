package routes

import (
	"github.com/gin-gonic/gin"
	teamConstructor "github.com/phn00dev/go-task-manager/internal/domain/team/constructor"
	todoConstructor "github.com/phn00dev/go-task-manager/internal/domain/todo/constructor"

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
		teamRoute := v1a.Group("/teams")
		{
			teamRoute.GET("/", teamConstructor.TeamHandler.GetAll)
			teamRoute.GET("/:teamID", teamConstructor.TeamHandler.GetOne)
			teamRoute.POST("/create", teamConstructor.TeamHandler.Create)
			teamRoute.PUT("/:teamID", teamConstructor.TeamHandler.Update)
			teamRoute.DELETE("/:teamID", teamConstructor.TeamHandler.Delete)
		}
		todoRoute := v1a.Group("/todos")
		{
			todoRoute.GET("/", todoConstructor.TodoHandler.GetAll)
			todoRoute.GET("/:todoID", todoConstructor.TodoHandler.GetOne)
			todoRoute.POST("/create", todoConstructor.TodoHandler.Create)
			todoRoute.PUT("/:todoID", todoConstructor.TodoHandler.Update)
			todoRoute.DELETE("/:todoID", todoConstructor.TodoHandler.Delete)
		}
	}
}
