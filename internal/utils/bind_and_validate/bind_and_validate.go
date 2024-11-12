package bindandvalidate

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-task-manager/internal/utils/response"
)

func BindAndValidate(ctx *gin.Context, request any) bool {
	if err := ctx.Bind(request); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return false
	}
	return true
}
