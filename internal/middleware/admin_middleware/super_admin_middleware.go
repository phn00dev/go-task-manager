package adminMiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-task-manager/internal/utils/response"
	admintoken "github.com/phn00dev/go-task-manager/pkg/jwt_token/adminToken"
	"net/http"
	"strings"
)

func SuperAdminMiddleware() gin.HandlerFunc {
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
		claims, err := admintoken.ValidateAdminToken(tokenString)
		if err != nil {
			response.Error(ctx, http.StatusUnauthorized, "error", "Invalid token")
			ctx.Abort()
		}
		if claims.AdminRole != "super_admin" {
			response.Error(ctx, http.StatusForbidden, "error", "Super admin access required")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Set("admin_id", claims.AdminID)
		ctx.Set("admin_role", claims.AdminRole)
		ctx.Next()
	}
}
