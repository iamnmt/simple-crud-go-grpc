package errors

type CustomError struct {
	Msg string
}

func (e CustomError) Error() string {
	return e.Msg
}
