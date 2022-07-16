package services

import (
	"encoding/json"
	"io/ioutil"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models"
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

func (svr *ServiceConsumerAdapter) GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (*models.ServiceIncomeResponse, error) {
	var serviceIncomeResponse = models.ServiceIncomeResponse{}
	url := viper.GetString("endpoint.strapi-information-gateway.host") + viper.GetString("endpoint.strapi-information-gateway.path")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return nil, err
	} else {
		err = nil
	}
	response, err := svr.UtilsPort.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return nil, err
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		serviceIncomeResponse.Error = err.Error()
		return nil, err
	} else {
		err = nil
	}
	return &serviceIncomeResponse, nil
}
