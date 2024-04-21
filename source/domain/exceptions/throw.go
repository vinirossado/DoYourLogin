package exceptions

func ThrowNotFoundException(message string, data any) {
	panic(NotFoundException(message, data))
}

func ThrowBadRequestException(message string, data any) {
	panic(BadRequestException(message, data))
}

func ThrowInternalServerException(message string, data any) {
	panic(InternalServerException(message, data))
}
