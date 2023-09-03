package ui

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/middleware"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server/util"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/static"
	"html/template"
	"net/http"
	"time"
)

func init() {
	uiTemplate = template.Must(template.New("name").Parse(string(static.FileWebIndexHTML)))
}

var uiTemplate *template.Template

func Index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	includeOffline := req.URL.Query().Get("include-offline") != "false"

	client := middleware.GetFritzBoxClient(req)
	netDevicesRes, err := fritzbox_requests.GetNetDevices(client)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Fritz!Box call failed")
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "public, max-age=60")
	w.Header().Set("Expires", time.Now().Add(1*time.Minute).Format(http.TimeFormat))
	w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

	err = uiTemplate.Execute(w, map[string]interface{}{
		"mapping": fritzbox_requests.MapToOnlineUsers(netDevicesRes, includeOffline),
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Rendering failed ("+err.Error()+"), please check template")
	}
}
