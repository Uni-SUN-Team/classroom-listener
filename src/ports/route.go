package ports

import "github.com/gin-gonic/gin"

type RouteConsumer interface {
	Consumer(r *gin.RouterGroup)
}
