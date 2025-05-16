package rest

import (
	"io"
	"locket-clone/backend/pkg/model"
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"
	"mime/multipart"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LocketController struct {
	locketAddingService  adding.LocketService
	locketListingService listing.LocketService
	userListingService   listing.UserService
}

const (
	DEFAULT_LOCKET_LIMIT = 50
)

func (controller *LocketController) GetFeed(ctx *gin.Context) {
	usernameVal, isUsernameSet := ctx.Get(AUTH_MIDDLEWARE_USERNAME_KEY)
	username, ok := usernameVal.(string)
	if !ok || !isUsernameSet {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
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
}

func (controller *LocketController) GetUserLockets(ctx *gin.Context) {
	username := ctx.Param("username")
	authUsernameReq, ok := ctx.Get(AUTH_MIDDLEWARE_USERNAME_KEY)
	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	authUsername, ok := authUsernameReq.(string)
	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

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
		startTimeUnix, err2 := strconv.ParseInt(startTimeString, 10, 64)
		if err2 != nil {
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
}

func (controller *LocketController) AddLocket(ctx *gin.Context) {
	usernameVal, isUsernameSet := ctx.Get(AUTH_MIDDLEWARE_USERNAME_KEY)
	username, ok := usernameVal.(string)
	if !ok || !isUsernameSet {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	type LocketPayload struct {
		Image   *multipart.FileHeader `json:"-" form:"image" binding:"required"`
		Caption string                `json:"caption" form:"caption" binding:"required"`
	}

	locket := LocketPayload{}
	if err := ctx.ShouldBind(&locket); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	fileType := locket.Image.Header.Get("Content-Type")
	if len(fileType) == 0 {
		ctx.JSON(400, gin.H{"error": "Invalid file type"})
		return
	}

	imageData, err := locket.Image.Open()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Cannot process image"})
		return
	}
	defer imageData.Close()

	imageBytes, err := io.ReadAll(imageData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Cannot process image"})
		return
	}

	payload := adding.LocketPayload{
		Type:     model.LocketType(fileType),
		Image:    imageBytes,
		Caption:  locket.Caption,
		Username: username,
	}
	if locket, err := controller.locketAddingService.AddLocket(payload); err != nil {
		ctx.JSON(500, gin.H{"error": "Error adding locket"})
		return
	} else {
		ctx.JSON(200, gin.H{
			"data": locket,
		})
		return
	}
}
