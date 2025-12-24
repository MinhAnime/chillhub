package error


type AppError struct {
	Status  int    // HTTP status
	Code    string // optional: error code (BUSINESS_xxx)
	Message string // message trả cho client
	Err     error  // lỗi gốc (dev trace)
}

func (e *AppError) Error() string {
	return e.Message
}

func New(status int, message string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
	}
}

func Wrap(err error, status int, message string) *AppError {
	return &AppError{
		Status:  status,
		Message: message,
		Err:     err,
	}
}