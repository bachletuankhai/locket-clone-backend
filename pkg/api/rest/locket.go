package rest

import "github.com/gin-gonic/gin"

// Fetch for latest 50 lockets, but need to refetch if specific user is chosen
// Fetch latest 10 lockets of each friend



func RegisterLocketHandler(group *gin.RouterGroup) {
	group.GET("/users/:id/feed")

}