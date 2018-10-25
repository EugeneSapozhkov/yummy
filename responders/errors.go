package responders

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func Error(w http.ResponseWriter, message string, status int) {
	response, _ := json.Marshal(ErrorResponse{Error: message, Status: status})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
