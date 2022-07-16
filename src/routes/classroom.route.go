package routes

import (
	"unisun/api/classroom-listener/src/ports"

	"github.com/gin-gonic/gin"
)

type RouteConsumerAdapter struct {
	Controller ports.ControllerConsumer
}

func NewRouteConsumerAdapter(controller ports.ControllerConsumer) *RouteConsumerAdapter {
	return &RouteConsumerAdapter{
		Controller: controller,
	}
}

func (srv *RouteConsumerAdapter) Consumer(g *gin.RouterGroup) {
	g.GET("/class-rooms", srv.Controller.ClassRooms)
	g.GET("/class-rooms/:id", srv.Controller.ClassRoomById)
	g.GET("/class-rooms/slug/:slug", srv.Controller.ClassRoomBySlug)
}
