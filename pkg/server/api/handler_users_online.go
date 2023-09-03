package api

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/middleware"
	"net/http"
)

func UsersOnline(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	client := middleware.GetFritzBoxClient(req)
	netDevicesRes, err := fritzbox_requests.GetNetDevices(client)
	if err != nil {
		server.SendError(w, http.StatusInternalServerError, "Fritz!Box call failed")
	}

	server.SendJson(w, fritzbox_requests.MapToOnlineUsers(netDevicesRes))
}
