package price

type ResponseSuccess struct {
	Data struct {
		Id           int     `json:"id"`
		ClassRoomId  int     `json:"classRoomId"`
		RegularPrice float64 `json:"regularPrice"`
		SpecialPrice float64 `json:"specialPrice"`
		Advisors     string  `json:"advisors"`
		Categories   string  `json:"categories"`
	} `json:"data"`
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
