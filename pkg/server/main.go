package server

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/api"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/middleware"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/ui"
	"net/http"
	"strconv"
	"time"
)

// Start the HTTP server
func Start(config *config.AppConfig, client *fritzbox_requests.FritzBoxClientWithRefresh) error {
	registerRoute := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		http.HandleFunc(pattern, middleware.Auth(middleware.Log(middleware.Context(config, client, handler))))
	}

	registerRoute("/", ui.Index)
	registerRoute("/api/users", api.UserList)
	registerRoute("/api/users/online", api.UsersOnline)
	registerRoute("/api/users/all", api.UsersAll)

	listen := ":" + strconv.Itoa(config.ServerPort)
	log.Print(log.CompServer, "Starting server on :"+listen+" ...")

	srv := http.Server{
		Addr:         listen,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}
