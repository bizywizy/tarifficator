package util

import (
	"net/http"
)

type ErrorDetail map[string][]string

type ErrorResponseBody struct {
	Message string      `json:"message"`
	Detail  ErrorDetail `json:"detail"`
}

func ResponseError(w http.ResponseWriter, status int, message string, detail ErrorDetail) {
	body := ErrorResponseBody{
		Message: message,
		Detail:  detail,
	}
	ResponseJson(w, status, body)
}
