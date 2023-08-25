package server

import (
	"context"
	"fmt"
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net/http"
)

const contextKeyConfig = "config"
const contextKeyFritzBoxClient = "fritzboxClient"

func getFritzBoxClient(r *http.Request) *fritzbox.Client {
	return r.Context().Value(contextKeyFritzBoxClient).(*fritzbox.Client)
}

func getConfig(r *http.Request) *config.AppConfig {
	return r.Context().Value(contextKeyConfig).(*config.AppConfig)
}

func ContextMiddleware(config *config.AppConfig, client *fritzbox.Client, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, originalRequest *http.Request) {
		config_context := context.WithValue(originalRequest.Context(), contextKeyConfig, config)
		client_context := context.WithValue(config_context, contextKeyFritzBoxClient, client)
		enhancedRequest := originalRequest.WithContext(client_context)

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		handler(w, enhancedRequest)
	}
}
