package api

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/util"
	"net/http"
)

func UserList(w http.ResponseWriter, req *http.Request) {
	deviceMapping := config.Get().DeviceNameMapping
	users := make([]string, 0)
	for user, _ := range deviceMapping {
		users = append(users, user)
	}
	util.SendJson(w, users)
}
