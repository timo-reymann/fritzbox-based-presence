package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
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
	if !ok || password != config.Get().AuthPassword {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false, true
	}

	return true, false
}

var authMapping = map[string]func(http.ResponseWriter, *http.Request) (bool, bool){
	"ip_range":         isAllowedIP,
	"www_authenticate": interceptWwwAuthenticate,
}

// Auth provides the authentication middleware and evaluates all the ones given in the order config, as soon
// as something passes the request is processed
func Auth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		order := config.Get().AuthMiddlewareOrder

		if len(order) == 0 {
			log.Print(log.CompAuth, "No middleware specified, skipping auth")
			handler(w, req)
			return
		}

		for _, name := range order {
			callback, ok := authMapping[name]
			log.Print(log.CompAuth, "Testing authentication "+name)
			if !ok {
				continue
			}

			authenticated, abort := callback(w, req)
			if authenticated {
				log.Print(log.CompAuth, "Authenticated using "+name)
				handler(w, req)
				return
			}

			if abort {
				log.Print(log.CompAuth, "Authenticated failed using "+name+", aborting")
				return
			}
		}

		log.Print(log.CompAuth, "Access denied")
		w.WriteHeader(http.StatusForbidden)
	}
}
