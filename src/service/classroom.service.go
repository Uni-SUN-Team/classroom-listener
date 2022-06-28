package service

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/model"
	"unisun/api/classroom-listener/src/ports"

	"github.com/spf13/viper"
)

type ServiceConsumerAdapter struct {
	UtilsPort ports.UtilsHTTPRequest
}

func New(utilsPort ports.UtilsHTTPRequest) *ServiceConsumerAdapter {
	return &ServiceConsumerAdapter{
		UtilsPort: utilsPort,
	}
}

func (svr *ServiceConsumerAdapter) GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (*model.ServiceIncomeResponse, error) {
	var serviceIncomeResponse = model.ServiceIncomeResponse{}
	url := viper.GetString("endpoint.strapi-information-gateway.host") + viper.GetString("endpoint.strapi-information-gateway.path")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
	}
	response, err := svr.UtilsPort.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return &serviceIncomeResponse, err
	} else {
		err = nil
	}
	return &serviceIncomeResponse, nil
}

func (srv *ServiceConsumerAdapter) GetAdvisorInfomation(id int64) (string, error) {
	_id := strconv.Itoa(int(id))
	url := strings.Join([]string{viper.GetString("endpoint.advisor.host"), viper.GetString("endpoint.advisor.path"), _id}, "")
	response, err := srv.UtilsPort.HTTPRequest(url, constants.GET, nil)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}

func (srv *ServiceConsumerAdapter) GetCoursesInformation(id int64) (string, error) {
	_id := strconv.Itoa(int(id))
	url := strings.Join([]string{viper.GetString("endpoint.courses.host"), viper.GetString("endpoint.courses.path"), _id}, "")
	response, err := srv.UtilsPort.HTTPRequest(url, constants.GET, nil)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}
