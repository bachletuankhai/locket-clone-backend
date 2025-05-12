package rest

import (
	"locket-clone/backend/pkg/service/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const AUTH_MIDDLEWARE_USERNAME_HEADER = "x-locket-username"

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if len(bearerToken) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	reqToken := strings.Split(bearerToken, " ")[1]
	if len(reqToken) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := auth.ParseToken(reqToken)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Request.Header.Set(AUTH_MIDDLEWARE_USERNAME_HEADER, claims.Username)
	ctx.Next()
}