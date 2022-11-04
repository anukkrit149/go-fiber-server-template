package structs

type Response struct {
	Success bool        `json:"success" schema:"success"`
	Data    interface{} `json:"data" schema:"data"`
}

type ResponseError struct {
	Success   bool        `json:"success" schema:"success"`
	ErrorCode string      `json:"error_code" schema:"error_code"`
	Error     interface{} `json:"errors" schema:"errors"`
}

type NameResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}
