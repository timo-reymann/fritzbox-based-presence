package server

import (
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net/http"
	"strconv"
)

// Start the HTTP server
func Start(config *config.AppConfig, client *fritzbox.Client) error {
	http.HandleFunc("/", LogMiddleware(ContextMiddleware(config, client, UIHandler)))
	http.HandleFunc("/api/users-online", LogMiddleware(ContextMiddleware(config, client, UsersOnlineHandler)))
	listen := "0.0.0.0:" + strconv.Itoa(config.ServerPort)
	println("Starting server on :" + listen + " ...")
	return http.ListenAndServe(listen, nil)
}
