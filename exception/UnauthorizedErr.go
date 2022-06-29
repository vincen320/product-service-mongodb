package exception

type UnauthorizedError struct {
	ErrString string
}

func NewUnauthorizedErr(e string) *UnauthorizedError {
	return &UnauthorizedError{
		ErrString: e,
	}
}

func (uae UnauthorizedError) Error() string {
	return uae.ErrString
}
