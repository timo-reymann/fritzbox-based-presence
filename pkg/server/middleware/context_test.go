package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContext(t *testing.T) {
	middleware := Context(&config.AppConfig{}, &fritzbox_requests.FritzBoxClientWithRefresh{}, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	middleware(httptest.NewRecorder(), req)

	if cfg := GetConfig(req); cfg == nil {
		t.Error("Expected config to be added to context")
	}

	if client := GetFritzBoxClient(req); client == nil {
		t.Error("Expected fritzbox client to be added to context")
	}
}
