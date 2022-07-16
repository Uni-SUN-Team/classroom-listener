package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/ports"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ControllerClassroomAdapter struct {
	Service ports.ServiceConsumer
}

func NewControllerClassroomAdapter(service ports.ServiceConsumer) *ControllerClassroomAdapter {
	return &ControllerClassroomAdapter{
		Service: service,
	}
}

// ClassRooms godoc
// @summary      	ClassRooms Listener
// @description  	ClassRooms Listener for the service
// @id           	ClassRoomsHandler
// @tags         	Feedback
// @accept       	json
// @produce      	json
// @success      	200    	{object}    addfeedback.ResponseClassRoomsSuccess    	"StatusOK"
// @failure      	401    	{object}   	addfeedback.ResponseFail    	"StatusUnauthorized"
// @failure     	422    	{object}   	addfeedback.ResponseFail   		"StatusUnprocessableEntity"
// @failure     	500    	{object}   	addfeedback.ResponseFail    	"StatusInternalServerError"
// @router       	/classroom-listener/api/v1/class-rooms [get]
func (srv *ControllerClassroomAdapter) ClassRooms(c *gin.Context) {
	payloadRequest := models.ServiceIncomeRequest{}
	payloadRequest.Method = constants.GET
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.path"), viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.query.value")}, "")
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		strings.Join([]string{payloadRequest.Path, "&", strings.Trim(query, "?")}, "")
	}
	payloadResponseFail := models.ResponseFail{}
	data, err := srv.Service.GetInformationFormStrapi(payloadRequest)
	if err != nil {
		payloadResponseFail.Error.Name = "InternalServerError"
		payloadResponseFail.Error.Status = http.StatusInternalServerError
		payloadResponseFail.Error.Message = err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, payloadResponseFail)
		return
	}

	if !data.Status {
		payloadResponseFail.Error.Name = "NotFound"
		payloadResponseFail.Error.Status = http.StatusNotFound
		payloadResponseFail.Error.Message = data.Error
		c.AbortWithStatusJSON(http.StatusNotFound, payloadResponseFail)
		return
	}

	payloadClassrooms := models.ResponseClassRoomsSuccess{}
	if err := json.Unmarshal([]byte(data.Payload), &payloadClassrooms); err != nil {
		payloadResponseFail.Error.Name = "UnprocessableEntity"
		payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
		payloadResponseFail.Error.Message = err.Error()
		payloadResponseFail.Error.Details = err
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
		return
	}

	parallelization := len(payloadClassrooms.Data)
	var wg sync.WaitGroup
	wg.Add(parallelization)
	for _, classRoom := range payloadClassrooms.Data {
		go func(Data *models.ClassRoom) {
			
		}(&classRoom)
	}

	c.AbortWithStatusJSON(http.StatusOK, payloadClassrooms)
}

func (srv *ControllerClassroomAdapter) ClassRoomById(c *gin.Context) {

}

func (srv *ControllerClassroomAdapter) ClassRoomBySlug(c *gin.Context) {

}
