package user_api

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type ResponseBody struct {
	Result  interface{} `json:"result,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    interface{} `json:"meta"`
}

type MetaInfo struct {
	Status int `json:"status"`
}

type MetaPagination struct {
	Status int `json:"status"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func NewMetaPagination(status, limit, offset, total int) MetaPagination {
	return MetaPagination{
		Status: status,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

type ErrorBody struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func OK(w http.ResponseWriter, data interface{}, msg string) {
	response := ResponseBody{
		Result:  data,
		Message: msg,
		Meta:    MetaInfo{http.StatusOK},
	}
	write(w, response, http.StatusOK)
}

func OKWithMeta(w http.ResponseWriter, data interface{}, msg string, meta interface{}) {
	response := ResponseBody{
		Result:  data,
		Message: msg,
		Meta:    meta,
	}
	write(w, response, http.StatusOK)
}

func Created(w http.ResponseWriter, data interface{}, msg string) {
	response := ResponseBody{
		Result:  data,
		Message: msg,
		Meta:    MetaInfo{http.StatusCreated},
	}
	write(w, response, http.StatusCreated)
}

func Error(w http.ResponseWriter, err error) {
	var errBody ErrorBody
	status := http.StatusInternalServerError

	switch origin := errors.Cause(err).(type) {
	case UserApiError:
		errBody = ErrorBody{
			Message: origin.Message,
			Code:    origin.ErrorCode,
		}
		status = origin.HTTPStatus
	default:
		errBody = ErrorBody{
			Message: "Internal Server Error",
			Code:    999,
		}
	}

	response := ResponseBody{
		Error: &errBody,
		Meta:  MetaInfo{status},
	}
	write(w, response, status)
}

func write(w http.ResponseWriter, result interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(result)
}
