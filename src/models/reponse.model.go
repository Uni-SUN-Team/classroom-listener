package models

type ResponseClassRoomsSuccess struct {
	Data []ClassRoom `json:"data"`
	Meta struct {
		Pagination paginationContent `json:"pagination"`
	} `json:"meta"`
}

type ResponseClassRoomSuccess struct {
	Data ClassRoom `json:"data"`
	Meta struct {
		Pagination paginationContent `json:"pagination"`
	} `json:"meta"`
}

type paginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}

type ResponseFail struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

type error struct {
	Status  int64       `json:"status"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
