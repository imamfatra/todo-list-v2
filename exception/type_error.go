package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

type UnauthorizedErr struct {
	Error string
}

func NewUnauthorizedErr(err string) UnauthorizedErr {
	return UnauthorizedErr{Error: err}
}
