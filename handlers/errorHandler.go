package handlers

import (
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type APIError interface {
	// APIError returns an HTTP status code and an API-safe error message.
	APIError() (int, string)
}

type restApiError struct {
	status int
	msg    string
}

func (e restApiError) Error() string {
	return e.msg
}

func (e restApiError) APIError() (int, string) {
	return e.status, e.msg
}

var (
	ErrAuth                = restApiError{status: http.StatusUnauthorized, msg: "invalid token"}
	ErrNotFound            = restApiError{status: http.StatusNotFound, msg: "not found"}
	ErrDuplicate           = restApiError{status: http.StatusBadRequest, msg: "duplicate"}
	ErrBadRquest           = restApiError{status: http.StatusBadRequest, msg: "bad request"}
	ErrInternalServerError = restApiError{status: http.StatusInternalServerError, msg: "internal server error"}
	ErrMissingKey          = restApiError{status: http.StatusBadRequest, msg: "Url Param 'key' is missing"}
)

func ErrorHandler(logger *zap.Logger, w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrBadRquest):
		http.Error(w, "failed to decode body", http.StatusBadRequest)
		logger.Error("failed to decode body", zap.Int("status", http.StatusBadRequest), zap.String("description", "Bad Request"))
	case errors.Is(err, ErrInternalServerError):
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		logger.Error("failed to encode JSON", zap.Int("status", http.StatusInternalServerError), zap.String("description", "internal server error"))
	case errors.Is(err, ErrMissingKey):
		http.Error(w, "Url Param 'key' is missing", http.StatusBadRequest)
		logger.Error("Url Param 'key' is missing", zap.Int("status", http.StatusBadRequest), zap.String("description", "Bad Request"))
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
		logger.Error("internal server error")
	}
}
