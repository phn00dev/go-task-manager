package app

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/setup/routes"
	"github.com/phn00dev/go-task-manager/pkg/config"
)

func NewApp(config *config.Config) (httpServer *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	httpServer = gin.New()

	httpServer.Use(gin.Logger())
	httpServer.Use(gin.Recovery())
	httpServer.Use(func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(500, gin.H{
					"status":  500,
					"message": "Näsazlyk ýüze çykdy, Sonrak synanysyn!!!",
				})
			}
		}()
		ctx.Next()
	})
	httpServer.Use(func(c *gin.Context) {
		c.Set("AppName", config.HttpServer.AppName)
		c.Set("ServerHeader", config.HttpServer.AppHeader)
	})

	// routes
	routes.AdminRoutes(httpServer)
	routes.UserRoutes(httpServer)
	return httpServer
}
