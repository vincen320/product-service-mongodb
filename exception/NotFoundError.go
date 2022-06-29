package exception

type NotFoundError struct {
	ErrString string
}

func NewNotFoundErr(e string) *NotFoundError {
	return &NotFoundError{
		ErrString: e,
	}
}

func (nfe NotFoundError) Error() string {
	return nfe.ErrString
}
