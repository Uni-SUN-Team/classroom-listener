package components

import (
	"encoding/json"
	"log"
	"unisun/api/classroom-listener/src/models/advisor"
	"unisun/api/classroom-listener/src/ports"
)

type ServiceConsumerAdap struct {
	Service ports.ServiceConsumer
}

func NewServiceConsumerAdap(service ports.ServiceConsumer) *ServiceConsumerAdap {
	return &ServiceConsumerAdap{
		Service: service,
	}
}

func (srv *ServiceConsumerAdap) MappingAdvisor(value []advisor.AdvisorData) ([]advisor.AdvisorData, error) {
	resultAdvisors := []advisor.AdvisorData{}
	for _, a := range value {
		advisorForm := advisor.ResponseAdvisor{}
		result, err := srv.Service.GetAdvisorInfomation(a.Id)
		if err != nil {
			log.Panic(err)
		}
		err = json.Unmarshal([]byte(result), &advisorForm)
		if err != nil {
			log.Panic(err)
		}
		resultAdvisors = append(resultAdvisors, advisorForm.Data)
	}
	return resultAdvisors, nil
}
