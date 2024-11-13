package userMiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-task-manager/internal/utils/response"
	usertoken "github.com/phn00dev/go-task-manager/pkg/jwt_token/userToken"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(ctx, http.StatusUnauthorized, "error", "Authorization header required")
			ctx.AbortWithStatus(401)
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.Error(ctx, http.StatusUnauthorized, "error", "Bearer token required")
			ctx.Abort()
			return
		}
		claims, err := usertoken.ValidateUserToken(tokenString)
		if err != nil {
			response.Error(ctx, http.StatusUnauthorized, "error", "Invalid token")
			ctx.Abort()
		}
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
