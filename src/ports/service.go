package ports

import (
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/models/advisor"
)

type ServiceConsumer interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (*models.ServiceIncomeResponse, error)
}

type ServiceAdvisor interface {
	GetAdivisor(id string) (*advisor.ResponseAdvisors, error)
}
