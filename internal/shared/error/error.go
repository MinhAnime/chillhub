package error

import (
	"errors"
	"net/http"
)

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

// WithErr tạo ra một bản sao của AppError nhưng có kèm thêm lỗi gốc (original error)
// Cách này giúp giữ nguyên Status và Message của lỗi mẫu
// Nếu truyền thêm tham số msg (tối đa 1), nó sẽ ghi đè Message của lỗi mẫu.
func (e *AppError) WithErr(err error, msg ...string) *AppError {
	newMsg := e.Message
	if len(msg) > 0 && msg[0] != "" {
		newMsg = msg[0]
	}

	return &AppError{
		Status:  e.Status,
		Code:    e.Code,
		Message: newMsg,
		Err:     err,
	}
}

func GetStatus(err error) int {
	if err == nil {
		return 200
	}
	var appErr *AppError
	// Sử dụng errors.As bên trong để code bên ngoài gọn hơn
	if errors.As(err, &appErr) {
		return appErr.Status
	}
	return http.StatusInternalServerError
}