package models

import "unisun/api/classroom-listener/src/models/course"

type Courses struct {
	Id          int64  `json:"id"`
	NameSubject string `json:"name_subject"`
	Courses     []data `json:"courses"`
}

type data struct {
	Id      int64             `json:"id"`
	Preview bool              `json:"preview"`
	Course  course.CourseData `json:"course"`
}
