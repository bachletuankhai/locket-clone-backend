package rest

import (
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserAddingService adding.UserService
	AuthService       auth.AuthService
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user adding.UserPayload
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := uc.UserAddingService.AddUser(user); err != nil {
		ctx.JSON(500, gin.H{"error": "Error adding user"})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (uc *UserController) Login(ctx *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	token, err := uc.AuthService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
