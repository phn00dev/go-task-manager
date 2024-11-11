package routes

import (
	"github.com/gin-gonic/gin"

	authConstructor "github.com/phn00dev/go-task-manager/internal/domain/auth/constructor"
	userConstructor "github.com/phn00dev/go-task-manager/internal/domain/user/constructor"
)

func UserRoutes(route *gin.Engine) {
	v1 := route.Group("/v1/api")
	{
		// auth routes
		authRoute := v1.Group("/auth")
		{
			authRoute.POST("/register", authConstructor.AuthHandler.Register)
			authRoute.POST("/register", authConstructor.AuthHandler.Login)
		}
		// user profile routes
		userRoute := v1.Group("/user")
		{
			userRoute.GET("/profile", userConstructor.UserHandler.GetUserProfile)
			userRoute.GET("/update-profile", userConstructor.UserHandler.UpdateData)
			userRoute.GET("/update-password", userConstructor.UserHandler.UpdatePassword)
			userRoute.GET("/delete-profile", userConstructor.UserHandler.DeleteProfile)
		}
	}
}