package view

// PingAPIResponse for GET
type PingAPIResponse struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

// Todo model
type Todo struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}

// TodoAPIResponse for DELETE
type TodoAPIResponse struct {
	Status string `json:"status"`
}
