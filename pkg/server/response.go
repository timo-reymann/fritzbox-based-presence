package server

import (
	"encoding/json"
	"net/http"
)

func sendJson(w http.ResponseWriter, response interface{}) {
	serialized, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(serialized)
}

func sendError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	sendJson(w, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}
