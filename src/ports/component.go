package ports

import "unisun/api/classroom-listener/src/models/advisor"

type ComponentMappingAdvisor interface {
	MappingAdvisor(value []advisor.AdvisorData) ([]advisor.AdvisorData, error)
}
