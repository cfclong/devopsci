package errors

import (
	"fmt"
	"net/http"
	"strings"
)

func OrmError1062(err error) bool {
	return strings.Contains(err.Error(), "Error 1062")
}

type Error struct {
	status  int
	code    string
	message string
	cause   error
}

func (this *Error) Error() string {
	return fmt.Sprintf("Error: %v, %v, %v, %v", this.status, this.code, this.message, this.cause)
}

func (this *Error) Status() int {
	return this.status
}

func (this *Error) Code() string {
	return this.code
}

func (this *Error) Message() string {
	return this.message
}

func (this *Error) Cause() error {
	return this.cause
}

func (this *Error) SetCode(code string) *Error {
	this.code = code
	return this
}

func (this *Error) SetMessage(format string, args ...interface{}) *Error {
	this.message = fmt.Sprintf(format, args...)
	return this
}

func (this *Error) SetCause(err error) *Error {
	this.cause = err
	return this
}

// Check following URL before add any new functions:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status

func NewBadRequest() *Error {
	return &Error{
		status:  http.StatusBadRequest,
		code:    "BadRequest",
		message: "bad request",
	}
}

func NewConflict() *Error {
	return &Error{
		status:  http.StatusConflict,
		code:    "Conflict",
		message: "conflict",
	}
}

func NewUnauthorized() *Error {
	return &Error{
		status:  http.StatusUnauthorized,
		code:    "Unauthorized",
		message: "unauthorized",
	}
}

func NewForbidden() *Error {
	return &Error{
		status:  http.StatusForbidden,
		code:    "Forbidden",
		message: "forbidden",
	}
}

func NewNotFound() *Error {
	return &Error{
		status:  http.StatusNotFound,
		code:    "NotFound",
		message: "not found",
	}
}

func NewMethodNotAllowed() *Error {
	return &Error{
		status:  http.StatusMethodNotAllowed,
		code:    "MethodNotAllowed",
		message: "method not allowed",
	}
}

func NewInternalServerError() *Error {
	return &Error{
		status:  http.StatusInternalServerError,
		code:    "InternalServerError",
		message: "internal server error",
	}
}
