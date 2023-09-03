package server

import (
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/api"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/middleware"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/ui"
	"net/http"
	"strconv"
)

// Start the HTTP server
func Start(config *config.AppConfig, client *fritzbox.Client) error {
	registerRoute := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		http.HandleFunc(pattern, middleware.Auth(middleware.Log(middleware.Context(config, client, handler))))
	}

	registerRoute("/", ui.Index)
	registerRoute("/api/users-online", api.UsersOnline)

	listen := "0.0.0.0:" + strconv.Itoa(config.ServerPort)
	println("Starting server on :" + listen + " ...")
	return http.ListenAndServe(listen, nil)
}
