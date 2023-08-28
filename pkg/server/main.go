package server

import (
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net/http"
	"strconv"
)

// Start the HTTP server
func Start(config *config.AppConfig, client *fritzbox.Client) error {
	registerRoute := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		http.HandleFunc(pattern, LogMiddleware(ContextMiddleware(config, client, handler)))
	}

	registerRoute("/", UIHandler)
	registerRoute("/api/users-online", UsersOnlineHandler)

	listen := "0.0.0.0:" + strconv.Itoa(config.ServerPort)
	println("Starting server on :" + listen + " ...")
	return http.ListenAndServe(listen, nil)
}
