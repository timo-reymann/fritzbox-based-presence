package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net"
	"net/http"
)

func isAllowedIP(_ http.ResponseWriter, req *http.Request) (bool, bool) {
	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return false, false
	}

	reqIp := net.ParseIP(host)

	if reqIp.IsLoopback() {
		return true, false
	}

	return config.AllowsIP(reqIp), false
}

func interceptWwwAuthenticate(w http.ResponseWriter, req *http.Request) (bool, bool) {
	_, password, ok := req.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false, true
	}

	if password == config.Get().AuthPassword {
		return true, true
	}

	return false, false
}

var authMapping = map[string]func(http.ResponseWriter, *http.Request) (bool, bool){
	"ip_range":         isAllowedIP,
	"www_authenticate": interceptWwwAuthenticate,
}

// Auth provides the authentication middleware and evaluates all the ones given in the order config, as soon
// as something passes the request is processed
func Auth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		for _, name := range config.Get().AuthMiddlewareOrder {
			callback, ok := authMapping[name]
			if ok {
				authenticated, abort := callback(w, req)
				if authenticated {
					println("[auth] Authenticated using " + name)
					handler(w, req)
					return
				} else if abort {
					println("[auth] Authenticated failed using " + name + ", aborting.")
					return
				}
			}
		}

		w.WriteHeader(http.StatusForbidden)
	}
}
