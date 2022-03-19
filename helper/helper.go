package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Total   int    `json:"total"`
}

func APIResponse(message string, code int, status string, total int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
		Total:   total,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	return response
}
