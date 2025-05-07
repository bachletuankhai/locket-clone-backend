package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}