package ports

import (
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/models/advisor"
	"unisun/api/classroom-listener/src/models/course"
	"unisun/api/classroom-listener/src/models/price"
)

type ServiceConsumer interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (*models.ServiceIncomeResponse, error)
}

type ServiceAdvisor interface {
	GetAdivisor(id string) (*advisor.ResponseAdvisor, error)
}

type ServiceCourse interface {
	GetCourse(id string) (*course.ResponseCourse, error)
}

type ServicePrice interface {
	GetClassRoomPrice(id string) (*price.ResponseSuccess, error)
}
