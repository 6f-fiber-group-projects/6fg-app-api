package libs

type ErrorWrapper struct {
	Hash    string
	Message string
}

func (e *ErrorWrapper) Error() *ErrorWrapper {
	return e
}

func FormatError(hash, msg string) *ErrorWrapper {
	return &ErrorWrapper{Hash: hash, Message: msg}
}
