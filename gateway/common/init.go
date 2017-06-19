package common

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ResponseError(w http.ResponseWriter, handle error, message string, status int) {
	errRs := errorResponse{
		Error:   handle.Error(),
		Message: message,
		Status:  status,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if j, err := json.Marshal(errRs); err == nil {
		w.Write(j)
	}
}
