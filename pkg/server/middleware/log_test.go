package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLog(t *testing.T) {
	middleware := Log(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middleware(httptest.NewRecorder(), httptest.NewRequest("GET", "/", strings.NewReader("")))
}
