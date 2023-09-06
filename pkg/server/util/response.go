package util

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// SendJson serializes the response
func SendJson(w http.ResponseWriter, response interface{}) {
	serialized, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}

// SendError serializes the error response
func SendError(w http.ResponseWriter, status int, message string) {
	println("[err] " + strconv.Itoa(status) + ": " + message)
	w.WriteHeader(status)
	SendJson(w, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}
