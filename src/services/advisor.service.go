package services

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unisun/api/classroom-listener/src/constants"
	"unisun/api/classroom-listener/src/models/advisor"
	"unisun/api/classroom-listener/src/ports"

	"github.com/spf13/viper"
)

type ServiceAdvisorAdapter struct {
	Utils ports.UtilsHTTPRequest
}

func NewServiceAdvisorAdapter(util ports.UtilsHTTPRequest) *ServiceAdvisorAdapter {
	return &ServiceAdvisorAdapter{
		Utils: util,
	}
}

func (srv *ServiceAdvisorAdapter) GetAdivisor(id string) (*advisor.ResponseAdvisors, error) {
	path := strings.Join([]string{viper.GetString("endpoint.advisor.host"), viper.GetString("endpoint.advisor.path")}, "")
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
	payloadAdvisor := advisor.ResponseAdvisors{}
	if err := json.Unmarshal([]byte(body), &payloadAdvisor); err != nil {
		return nil, err
	}
	return &payloadAdvisor, nil
}
