package server

import "net/http"

func LogMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		println("[request] " + req.Method + " " + req.URL.Path + " by " + req.RemoteAddr)
		handler(w, req)
	}
}
