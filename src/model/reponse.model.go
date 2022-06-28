package model

type ResponseSuccess struct {
	Data       []classRoom `json:"data"`
	Pagination pagination  `json:"pagination"`
}

type ResponseSingleSuccess struct {
	Data       classRoom  `json:"data"`
	Pagination pagination `json:"pagination"`
}

type pagination struct {
	Pagination paginationContent `json:"pagination"`
}

type paginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}
