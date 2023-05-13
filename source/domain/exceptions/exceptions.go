package exception

import "net/http"

type HttpException struct {
	StatusCode int    `example:"400"`
	Message    string `example:"Invalid path parameter"`
}

func (h *HttpException) Error() string {
	return h.Message
}

func NotFoundException(message string) *HttpException {
	return &HttpException{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func BadRequestException(message string) *HttpException {
	return &HttpException{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func InternalServerException(message string) *HttpException {
	return &HttpException{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
