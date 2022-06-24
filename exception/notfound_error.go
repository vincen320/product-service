package exception

type NotFoundError struct {
	ErrorString string
}

//Mau coba return yang pointer bisa tidak
//jawabannya bisa
func NewNotFoundError(errString string) *NotFoundError {
	return &NotFoundError{
		ErrorString: errString,
	}
}
func (b *NotFoundError) Error() string {
	return b.ErrorString
}
