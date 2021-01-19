package helper

//Response struct
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//Meta struct
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

//APIResponse for format API response
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}
