package advisor

type ResponseAdvisors struct {
	advisors
}

type ResponseAdvisor struct {
	advisor
}

type ResponseAdvisorFail struct {
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

type error struct {
	Status  int64       `json:"status"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
