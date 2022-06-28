package model

import "unisun/api/classroom-listener/src/model/course"

type courses struct {
	Id          int64  `json:"id"`
	NameSubject string `json:"name_subject"`
	Courses     []data `json:"courses"`
}

type data struct {
	Id      int64      `json:"id"`
	Preview bool       `json:"preview"`
	Course  CourseData `json:"course"`
}

type CourseData struct {
	course.CourseData
}
