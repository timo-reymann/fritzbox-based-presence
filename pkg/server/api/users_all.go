package api

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/middleware"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/util"
	"net/http"
)

func UsersAll(w http.ResponseWriter, req *http.Request) {
	client := middleware.GetFritzBoxClient(req)
	netDevicesRes, err := fritzbox_requests.GetNetDevices(client)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.SendJson(w, fritzbox_requests.MapToOnlineUsers(netDevicesRes, true))
}
