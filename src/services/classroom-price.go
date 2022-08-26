package services

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models/price"
	"unisun/api/classroom-listener/src/ports"

	"github.com/spf13/viper"
)

type ServiceClassRoomPrice struct {
	Utils ports.UtilsHTTPRequest
}

func NewServiceClassRoomPrice(util ports.UtilsHTTPRequest) *ServiceClassRoomPrice {
	return &ServiceClassRoomPrice{
		Utils: util,
	}
}

func (srv *ServiceClassRoomPrice) GetClassRoomPrice(id string) (*price.ResponseSuccess, error) {
	path := strings.Join([]string{viper.GetString("endpoint.price.host"), viper.GetString("endpoint.price.path"), id}, "")
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
	payloadPrice := price.ResponseSuccess{}
	if err := json.Unmarshal([]byte(body), &payloadPrice); err != nil {
		return nil, err
	}
	return &payloadPrice, nil
}
