package libs

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResSuccess(message string, data interface{}) *Response {
	return &Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ResError(message string) *Response {
	return &Response{
		Status:  "error",
		Message: message,
	}
}
