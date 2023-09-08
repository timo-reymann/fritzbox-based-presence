package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
	"net/http"
)

func Log(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		log.Print(log.CompServer, req.Method+" "+req.URL.Path+" by "+req.RemoteAddr)
		handler(w, req)
	}
}
