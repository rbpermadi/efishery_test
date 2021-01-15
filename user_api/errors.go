package user_api

import (
	"fmt"
	"net/http"
)

var (
	ErrNotFound = UserApiError{
		Message:    "Tidak ditemukan",
		ErrorCode:  404,
		HTTPStatus: http.StatusNotFound,
	}

	ErrInvalidParameter = UserApiError{
		Message:    "Parameter tidak valid",
		ErrorCode:  400,
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidCredentials = UserApiError{
		Message:    "Phone atau password salah",
		ErrorCode:  401,
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrUnauthorized = UserApiError{
		Message:    "Unauthorized",
		ErrorCode:  401,
		HTTPStatus: http.StatusUnauthorized,
	}
)

type UserApiError struct {
	Message    string `json:"message"`
	ErrorCode  int    `json:"error_code"`
	HTTPStatus int    `json:"-"`
}

func (e UserApiError) Error() string {
	return e.Message
}

func ValidationError(err error) UserApiError {
	return UserApiError{
		Message:    err.Error(),
		ErrorCode:  400,
		HTTPStatus: http.StatusBadRequest,
	}
}

func CustomValidationError(errmsg string, args ...interface{}) UserApiError {
	return UserApiError{
		Message:    fmt.Sprintf(errmsg, args...),
		ErrorCode:  400,
		HTTPStatus: http.StatusBadRequest,
	}
}
