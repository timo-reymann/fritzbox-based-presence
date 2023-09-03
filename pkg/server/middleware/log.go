package middleware

import (
	"net/http"
)

func Log(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		println("[request] " + req.Method + " " + req.URL.Path + " by " + req.RemoteAddr)
		handler(w, req)
	}
}
