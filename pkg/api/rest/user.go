package rest

import (
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/auth"
)

type UserController struct {
	UserAddingService adding.UserService
	AuthService       auth.AuthService
}
