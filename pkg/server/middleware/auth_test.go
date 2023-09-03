package middleware

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newRequest(remoteAddr string, addAuth bool) *http.Request {
	req := httptest.NewRequest("GET", "/", strings.NewReader(""))
	if remoteAddr == "" {
		req.RemoteAddr = "192.168.178.20:5555"
	} else {
		req.RemoteAddr = remoteAddr + ":5555"
	}

	if addAuth {
		req.Header.Set("Authorization", "Basic dGVzdDpmb28=")
	}

	return req
}

func TestAuth(t *testing.T) {
	middleware := Auth(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		orderConfig    string
		req            *http.Request
		expectedStatus int
	}{
		{
			"",
			newRequest("", false),
			http.StatusOK,
		},
		{
			"ip_range",
			newRequest("", false),
			http.StatusOK,
		},
		{
			"ip_range",
			newRequest("127.0.0.1", false),
			http.StatusOK,
		},
		{
			"ip_range",
			newRequest("10.0.0.1", false),
			http.StatusForbidden,
		},
		{
			"www_authenticate",
			newRequest("", false),
			http.StatusUnauthorized,
		},
		{
			"www_authenticate",
			newRequest("", true),
			http.StatusOK,
		},
		{
			"ip_range,www_authenticate",
			newRequest("1.1.1.1", true),
			http.StatusOK,
		},
		{
			"www_authenticate,ip_range",
			newRequest("", true),
			http.StatusOK,
		},
	}

	for _, tc := range testCases {
		_ = config.Read()

		config.Get().AuthPassword = "foo"
		if tc.orderConfig != "" {
			config.Get().AuthMiddlewareOrder = strings.Split(tc.orderConfig, ",")
		}

		ipRange := config.IpRangeDecoder{}
		_ = ipRange.Decode("192.168.178.0/24")
		config.Get().AuthIpRange = ipRange

		rec := httptest.NewRecorder()
		middleware(rec, tc.req)
		if rec.Code != tc.expectedStatus {
			t.Fatalf("Expected status code %d but got %d", tc.expectedStatus, rec.Code)
		}
	}
}
