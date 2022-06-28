package ports

import "github.com/gin-gonic/gin"

type ControllerConsumer interface {
	ClassRooms(c *gin.Context)
	ClassRoomById(c *gin.Context)
	ClassRoomBySlug(c *gin.Context)
}
