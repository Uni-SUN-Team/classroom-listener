package services

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models/course"
	"unisun/api/classroom-listener/src/ports"

	"github.com/spf13/viper"
)

type ServiceCourseAdapter struct {
	Utils ports.UtilsHTTPRequest
}

func NewServiceCourseAdapter(util ports.UtilsHTTPRequest) *ServiceCourseAdapter {
	return &ServiceCourseAdapter{
		Utils: util,
	}
}

func (srv *ServiceCourseAdapter) GetCourse(id string) (*course.ResponseCourse, error) {
	path := strings.Join([]string{viper.GetString("endpoint.courses.host"), viper.GetString("endpoint.courses.path"), id}, "")
	response, err := srv.Utils.HTTPRequest(path, constants.GET, nil)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	} else {
		err = nil
		defer response.Body.Close()
	}
	payloadCourse := course.ResponseCourse{}
	if err := json.Unmarshal([]byte(body), &payloadCourse); err != nil {
		return nil, err
	}
	return &payloadCourse, nil
}
