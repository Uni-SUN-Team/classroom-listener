package ports

import "unisun/api/classroom-listener/src/model"

type ServiceConsumer interface {
	GetInformationFormStrapi(payloadRequest model.ServiceIncomeRequest) (*model.ServiceIncomeResponse, error)
	GetAdvisorInfomation(id int64) (string, error)
	GetCoursesInformation(id int64) (string, error)
}
