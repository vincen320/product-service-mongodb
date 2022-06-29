package exception

type BadRequestError struct {
	ErrString string
}

func NewBadRequestErr(e string) *BadRequestError {
	return &BadRequestError{
		ErrString: e,
	}
}

func (bre BadRequestError) Error() string {
	return bre.ErrString
}
