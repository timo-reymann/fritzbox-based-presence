package util

import (
	"encoding/json"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
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
	log.Print(log.CompServer, "Sending error response: "+strconv.Itoa(status)+" "+message)
	w.WriteHeader(status)
	SendJson(w, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}
