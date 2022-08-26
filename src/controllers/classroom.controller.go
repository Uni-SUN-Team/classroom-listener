package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/ports"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ControllerClassroomAdapter struct {
	Service           ports.ServiceConsumer
	MapAdvisor        ports.ComponentMappingAdvisor
	MapCourse         ports.ComponentMappingCourse
	MapClassRoomPrice ports.ComponentMappingClassRoomPrice
}

func NewControllerClassroomAdapter(service ports.ServiceConsumer,
	mapAdvisor ports.ComponentMappingAdvisor,
	mapCourse ports.ComponentMappingCourse,
	mapClassRoomPrice ports.ComponentMappingClassRoomPrice) *ControllerClassroomAdapter {
	return &ControllerClassroomAdapter{
		Service:           service,
		MapAdvisor:        mapAdvisor,
		MapCourse:         mapCourse,
		MapClassRoomPrice: mapClassRoomPrice,
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
	resp := models.ResponseClassRoomsSuccess{}
	for _, classRoom := range payloadClassrooms.Data {
		if resultMapPrice, err := srv.MapClassRoomPrice.MappingClassRoomPrice(classRoom); err != nil {
			payloadResponseFail.Error.Name = "UnprocessableEntity"
			payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
			payloadResponseFail.Error.Message = err.Error()
			payloadResponseFail.Error.Details = err
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
			return
		} else {
			resultMapAdvisor, err := srv.MapAdvisor.MappingAdvisor(resultMapPrice.Advisors)
			if err != nil {
				payloadResponseFail.Error.Name = "UnprocessableEntity"
				payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
				payloadResponseFail.Error.Message = err.Error()
				payloadResponseFail.Error.Details = err
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
				return
			}
			resultMapPrice.Advisors = *resultMapAdvisor
			resp.Data = append(resp.Data, *resultMapPrice)
		}
	}
	resp.Meta = payloadClassrooms.Meta
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func (srv *ControllerClassroomAdapter) ClassRoomById(c *gin.Context) {
	id := c.Param("id")
	payloadRequest := models.ServiceIncomeRequest{}
	payloadRequest.Method = constants.GET
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.path"), "/", id, viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.query.value")}, "")
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
	payloadClassroom := models.ResponseClassRoomSuccess{}
	if err := json.Unmarshal([]byte(data.Payload), &payloadClassroom); err != nil {
		payloadResponseFail.Error.Name = "UnprocessableEntity"
		payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
		payloadResponseFail.Error.Message = err.Error()
		payloadResponseFail.Error.Details = err
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
		return
	}
	resp := models.ResponseClassRoomSuccess{}
	classRoom := payloadClassroom.Data
	if resultMapPrice, err := srv.MapClassRoomPrice.MappingClassRoomPrice(classRoom); err != nil {
		payloadResponseFail.Error.Name = "UnprocessableEntity"
		payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
		payloadResponseFail.Error.Message = err.Error()
		payloadResponseFail.Error.Details = err
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
		return
	} else {
		resultMapAdvisor, err := srv.MapAdvisor.MappingAdvisor(resultMapPrice.Advisors)
		if err != nil {
			payloadResponseFail.Error.Name = "UnprocessableEntity"
			payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
			payloadResponseFail.Error.Message = err.Error()
			payloadResponseFail.Error.Details = err
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
			return
		}
		resultMapPrice.Advisors = *resultMapAdvisor
		resp.Data = *resultMapPrice

		resultMapCourse, err := srv.MapCourse.MappingCourse(resultMapPrice.Courses)
		if err != nil {
			payloadResponseFail.Error.Name = "UnprocessableEntity"
			payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
			payloadResponseFail.Error.Message = err.Error()
			payloadResponseFail.Error.Details = err
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
			return
		}
		resultMapPrice.Courses = *resultMapCourse
	}

	resp.Meta = payloadClassroom.Meta
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func (srv *ControllerClassroomAdapter) ClassRoomBySlug(c *gin.Context) {
	slug := c.Param("slug")
	payloadRequest := models.ServiceIncomeRequest{}
	payloadRequest.Method = constants.GET
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.path"), viper.GetString("endpoint.strapi-information-gateway.mapping.class-rooms.query.value"), "&filters[slug][$eq]=", slug}, "")
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
	payloadClassroom := models.ResponseClassRoomSuccess{}
	payloadClassroom.Data = payloadClassrooms.Data[0]
	payloadClassroom.Meta = payloadClassrooms.Meta
	resp := models.ResponseClassRoomSuccess{}
	classRoom := payloadClassroom.Data
	if resultMapPrice, err := srv.MapClassRoomPrice.MappingClassRoomPrice(classRoom); err != nil {
		payloadResponseFail.Error.Name = "UnprocessableEntity"
		payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
		payloadResponseFail.Error.Message = err.Error()
		payloadResponseFail.Error.Details = err
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
		return
	} else {
		resultMapAdvisor, err := srv.MapAdvisor.MappingAdvisor(resultMapPrice.Advisors)
		if err != nil {
			payloadResponseFail.Error.Name = "UnprocessableEntity"
			payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
			payloadResponseFail.Error.Message = err.Error()
			payloadResponseFail.Error.Details = err
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
			return
		}
		resultMapPrice.Advisors = *resultMapAdvisor
		resp.Data = *resultMapPrice

		resultMapCourse, err := srv.MapCourse.MappingCourse(resultMapPrice.Courses)
		if err != nil {
			payloadResponseFail.Error.Name = "UnprocessableEntity"
			payloadResponseFail.Error.Status = http.StatusUnprocessableEntity
			payloadResponseFail.Error.Message = err.Error()
			payloadResponseFail.Error.Details = err
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, payloadResponseFail)
			return
		}
		resultMapPrice.Courses = *resultMapCourse
	}

	resp.Meta = payloadClassroom.Meta
	c.AbortWithStatusJSON(http.StatusOK, resp)
}
