package components

import (
	"log"
	"strconv"
	"unisun/api/classroom-listener/src/models/advisor"
	"unisun/api/classroom-listener/src/ports"
)

type MappingAdvisorAdap struct {
	Service ports.ServiceAdvisor
}

func NewMappingAdvisorAdap(service ports.ServiceAdvisor) *MappingAdvisorAdap {
	return &MappingAdvisorAdap{
		Service: service,
	}
}

func (srv *MappingAdvisorAdap) MappingAdvisor(value []advisor.AdvisorData) (*[]advisor.AdvisorData, error) {
	resultAdvisors := []advisor.AdvisorData{}
	for _, a := range value {
		result, err := srv.Service.GetAdivisor(strconv.FormatInt(a.Id, 10))
		if err != nil {
			log.Panic(err)
		}
		resultAdvisors = append(resultAdvisors, result.Data)
	}
	return &resultAdvisors, nil
}
