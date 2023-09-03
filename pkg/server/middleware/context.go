package middleware

import (
	"context"
	"fmt"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"net/http"
)

const contextKeyConfig = "config"
const contextKeyFritzBoxClient = "fritzboxClient"

// GetFritzBoxClient stored in the request
func GetFritzBoxClient(r *http.Request) *fritzbox_requests.FritzBoxClientWithRefresh {
	return r.Context().Value(contextKeyFritzBoxClient).(*fritzbox_requests.FritzBoxClientWithRefresh)
}

// GetConfig returns the config stored in the request
func GetConfig(r *http.Request) *config.AppConfig {
	return r.Context().Value(contextKeyConfig).(*config.AppConfig)
}

// Context provides config and Fritz!Box client to the request
func Context(config *config.AppConfig, client *fritzbox_requests.FritzBoxClientWithRefresh, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, originalRequest *http.Request) {
		config_context := context.WithValue(originalRequest.Context(), contextKeyConfig, config)
		client_context := context.WithValue(config_context, contextKeyFritzBoxClient, client)
		enhancedRequest := originalRequest.WithContext(client_context)
		*originalRequest = *enhancedRequest

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		handler(w, enhancedRequest)
	}
}
