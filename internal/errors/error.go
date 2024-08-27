package errors

type Error struct {
	code string
	err  string
}

func (e Error) Error() string {
	return e.err
}

func (e Error) Code() string {
	return e.code
}
