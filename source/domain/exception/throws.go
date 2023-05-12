package exception

func ThrowNotFoundException(message string) {
	panic(NotFoundException(message))
}

func ThrowBadRequestException(message string) {
	panic(BadRequestException(message))
}

func ThrowInternalServerException(message string) {
	panic(InternalServerException(message))
}
