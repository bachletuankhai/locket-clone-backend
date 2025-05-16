package rest

import (
	"locket-clone/backend/pkg/service/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const AUTH_MIDDLEWARE_USERNAME_KEY = "x-locket-username"

func NewAuthMiddleware(authService auth.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		claims, err := authService.ParseToken(reqToken)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set(AUTH_MIDDLEWARE_USERNAME_KEY, claims.Username)
		ctx.Next()
	}
}
