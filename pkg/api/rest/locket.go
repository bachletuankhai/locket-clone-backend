package rest

import (
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Fetch for latest 50 lockets, but need to refetch if specific user is chosen
// Fetch latest 10 lockets of each friend

type LocketController struct {
	locketAddingService  adding.LocketService
	locketListingService listing.LocketService
	userListingService   listing.UserService
}

const (
	DEFAULT_LOCKET_LIMIT = 50
)

func (controller *LocketController) RegisterLocketHandler(group *gin.RouterGroup) {
	group.GET("/feed", AuthMiddleware, func(ctx *gin.Context) {
		username := ctx.GetHeader(AUTH_MIDDLEWARE_USERNAME_HEADER)
		startTimeString, isStartTimePresent := ctx.GetQuery("startTime")
		var lockets []listing.Locket
		var err error

		var startTime time.Time
		if isStartTimePresent {
			startTimeUnix, err := strconv.ParseInt(startTimeString, 10, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid start time"})
				return
			}
			startTime = time.Unix(startTimeUnix, 0)
		}

		userIds, err := controller.userListingService.GetVisibleUserIds(username)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error fetching user"})
			return
		}
		if isStartTimePresent {
			lockets, err = controller.locketListingService.ListLocketsByUserIdsTime(userIds, startTime, DEFAULT_LOCKET_LIMIT)
		} else {
			lockets, err = controller.locketListingService.ListLatestLockets(userIds, DEFAULT_LOCKET_LIMIT)
		}
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error fetching lockets"})
			return
		}
		ctx.JSON(200, gin.H{
			"data": lockets,
		})
	})

	group.GET("/user/:username", AuthMiddleware, func(ctx *gin.Context) {
		username := ctx.Param("username")
		authUsername := ctx.GetHeader(AUTH_MIDDLEWARE_USERNAME_HEADER)

		if username != authUsername {
			user, err := controller.userListingService.GetUserByUsername(username)
			if err != nil {
				ctx.JSON(500, gin.H{"error": "Error fetching user"})
				return
			}
			visibleUserIds, err := controller.userListingService.GetVisibleUserIds(authUsername)
			if err != nil {
				ctx.JSON(500, gin.H{"error": "Error fetching user"})
				return
			}
			if !slices.Contains(visibleUserIds, user.ID) {
				ctx.JSON(403, gin.H{"error": "You are not allowed to view this user's lockets"})
				return
			}
		}

		startTimeString, isStartTimePresent := ctx.GetQuery("startTime")
		var lockets []listing.Locket
		var err error

		var startTime time.Time
		if isStartTimePresent {
			startTimeUnix, err := strconv.ParseInt(startTimeString, 10, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid start time"})
				return
			}
			startTime = time.Unix(startTimeUnix, 0)

			lockets, err = controller.locketListingService.ListUserLocketsByUsernameTime(username, startTime, DEFAULT_LOCKET_LIMIT)
		} else {
			lockets, err = controller.locketListingService.ListUserLocketsByUsername(username, DEFAULT_LOCKET_LIMIT)
		}
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Error fetching lockets"})
			return
		}
		ctx.JSON(200, gin.H{
			"data": lockets,
		})
	})

}
