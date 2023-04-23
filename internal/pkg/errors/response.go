package errors

import "fmt"

type ResponseError struct {
	Code    int    // 错误码
	Message string // 错误信息
}

func (r *ResponseError) Error() string {
	return r.Message
}

func newResponse(code int, message string, args ...interface{}) error {
	return &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
	}
}
func NewCustomError(code int, message string, args ...interface{}) error {
	return &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
	}
}
