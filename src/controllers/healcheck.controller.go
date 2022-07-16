package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerHealthCheckHandlerAdapter struct {
}

func NewControllerHealthCheckHandler() *ControllerHealthCheckHandlerAdapter {
	return &ControllerHealthCheckHandlerAdapter{}
}

// HealthCheckHandler godoc
// @summary      Health Check
// @description  Health checking for the service
// @id           HealthCheckHandler
// @produce      plain
// @response     200  {string}  string  "OK"
// @router       /feedback-gateway/api/v1/healcheck [get]
func (*ControllerHealthCheckHandlerAdapter) HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
