package rest

import (
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"

	"github.com/gin-gonic/gin"
)

// Fetch for latest 50 lockets, but need to refetch if specific user is chosen
// Fetch latest 10 lockets of each friend


type LocketController struct {
	locketAddingService adding.LocketService
	locketListingService listing.LocketService
}

func (controller *LocketController) RegisterLocketHandler(group *gin.RouterGroup) {
	group.GET("/feed", AuthMiddleware, func(ctx *gin.Context) {
		username := ctx.GetHeader(AUTH_MIDDLEWARE_USERNAME_HEADER)
		controller.locketListingService.ListUserLocketsByUsername(username, )
	})

}