package components

import (
	"log"
	"strconv"
	"unisun/api/classroom-listener/src/models"
	"unisun/api/classroom-listener/src/ports"
)

type MappingCourseAdap struct {
	Service ports.ServiceCourse
}

func NewMappingCourseAdap(service ports.ServiceCourse) *MappingCourseAdap {
	return &MappingCourseAdap{
		Service: service,
	}
}

func (srv *MappingCourseAdap) MappingCourse(value []models.Courses) (*[]models.Courses, error) {
	resultCourse := []models.Courses{}
	for _, courses := range value {
		for index, course := range courses.Courses {
			result, err := srv.Service.GetCourse(strconv.FormatInt(course.Course.Id, 10))
			if err != nil {
				log.Panic(err)
			}
			courses.Courses[index].Course = result.Data
		}
		resultCourse = append(resultCourse, courses)
	}
	return &resultCourse, nil
}
