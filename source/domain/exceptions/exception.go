package exceptions

import "net/http"

type HttpException struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func (h *HttpException) Error() string {
	return h.Message
}

func NotFoundException(message string, data any) *HttpException {
	return &HttpException{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Data:       data,
	}
}

func BadRequestException(message string, data any) *HttpException {
	return &HttpException{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Data:       data,
	}
}

func InternalServerException(message string, data any) *HttpException {
	return &HttpException{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Data:       data,
	}
}
