package server

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"net/http"
)

func UsersOnlineHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	client := getFritzBoxClient(req)
	netDevicesRes, err := fritzbox_requests.GetNetDevices(client)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Fritz!Box call failed")
	}

	sendJson(w, fritzbox_requests.MapToOnlineUsers(netDevicesRes))
}
