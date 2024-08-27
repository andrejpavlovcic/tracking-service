package errors

var (
	ErrAccountNotFound = Error{
		"not_found",
		"Requested account was not found",
	}
)
