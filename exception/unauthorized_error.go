package exception

type UnauthorizedError struct {
	ErrorString string
}

//Mau coba return yang pointer bisa tidak
//jawabannya bisa
func NewUnauthorizedError(errString string) *UnauthorizedError {
	return &UnauthorizedError{
		ErrorString: errString,
	}
}
func (b *UnauthorizedError) Error() string {
	return b.ErrorString
}
