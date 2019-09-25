package api

import (
	"fmt"
)

type StatusError interface {
	error
	StatusCode() int
	ErrorCode() int
	Message() string
	ErrorInfo() map[string]interface{}
}

func IsStatusError(err error) bool {
	_, ok := err.(StatusError)
	return ok
}

func HasStatusCode(err error, code int) bool {
	if statusError, ok := err.(StatusError); ok {
		return statusError.StatusCode() == code
	}

	return false
}

func HasErrorCode(err error, code int) bool {
	if statusError, ok := err.(StatusError); ok {
		return statusError.ErrorCode() == code
	}

	return false
}

type statusError struct {
	statusCode int
	errorCode  int
	message    string
	errorInfo  map[string]interface{}
}

func NewStatusError(statusCode, errorCode int, message string, errorInfo map[string]interface{}) StatusError {
	return &statusError{
		statusCode: statusCode,
		errorCode:  errorCode,
		message:    message,
		errorInfo:  errorInfo,
	}
}

func (e *statusError) StatusCode() int {
	return e.statusCode
}

func (e *statusError) ErrorCode() int {
	return e.errorCode
}

func (e *statusError) Message() string {
	return e.message
}

func (e *statusError) ErrorInfo() map[string]interface{} {
	return e.errorInfo
}

func (e *statusError) Error() string {
	return fmt.Sprintf(
		"StatusCode: %d, ErrorCode: %d, Message: %q, Info: %v",
		e.statusCode, e.errorCode, e.message, e.errorInfo,
	)
}
