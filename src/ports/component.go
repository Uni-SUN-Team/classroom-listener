package ports

import (
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/models/advisor"
)

type ComponentMappingAdvisor interface {
	MappingAdvisor(value []advisor.AdvisorData) (*[]advisor.AdvisorData, error)
}

type ComponentMappingCourse interface {
	MappingCourse(value []models.Courses) (*[]models.Courses, error)
}

type ComponentMappingClassRoomPrice interface {
	MappingClassRoomPrice(value models.ClassRoom) (*models.ClassRoom, error)
}
