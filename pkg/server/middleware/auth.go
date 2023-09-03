package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net"
	"net/http"
)

func isAllowedIP(req *http.Request) bool {
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return false
	}

	reqIp := net.ParseIP(host)

	if reqIp.IsLoopback() {
		return true
	}

	return config.AllowsIP(reqIp)
}

var authMapping = map[string]func(req *http.Request) bool{
	"ip_range": isAllowedIP,
}

// Auth provides the authentication middleware and evaluates all the ones given in the order config, as soon
// as something passes the request is processed
func Auth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, name := range config.Get().AuthMiddlewareOrder {
			callback, ok := authMapping[name]
			if ok && callback(req) {
				handler(w, req)
				return
			}
		}

		w.WriteHeader(http.StatusForbidden)
	}
}
