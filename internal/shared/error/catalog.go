package error

import "net/http"

var (
	BadRequest = func(msg string) *AppError {
		return New(http.StatusBadRequest, msg)
	}

	NotFound = func(msg string) *AppError {
		return New(http.StatusNotFound, msg)
	}

	Forbidden = func(msg string) *AppError {
		return New(http.StatusForbidden, msg)
	}

	Internal = func(err error) *AppError {
		return Wrap(err, http.StatusInternalServerError, "internal.error")
	}
)