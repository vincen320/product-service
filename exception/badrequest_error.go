package exception

type BadRequestError struct {
	ErrorString string
}

//Mau coba return yang pointer bisa tidak
//jawabannya bisa
func NewBadRequestError(errString string) *BadRequestError {
	return &BadRequestError{
		ErrorString: errString,
	}
}
func (b *BadRequestError) Error() string {
	return b.ErrorString
}
