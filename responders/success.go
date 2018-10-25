package responders

import (
	"encoding/json"
	"net/http"
)

type SuccessType struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Success(w http.ResponseWriter, message string, status int) {
	response, _ := json.Marshal(SuccessType{Message: message, Status: status})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
