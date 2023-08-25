package fritzbox_requests

import (
	"crypto/tls"
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"net/http"
	"net/url"
)

func createHttpClient(ignoreCertificates bool) *http.Client {
	httpClient := http.DefaultClient
	if ignoreCertificates {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		httpClient = &http.Client{Transport: tr}
	}
	return httpClient
}

// CreateAuthenticatedFritzBoxClient generates an fritzbox client based on given configuration.
// After parsing the url a login is attempted, if that fails an error is returned
func CreateAuthenticatedFritzBoxClient(config *config.AppConfig) (*fritzbox.Client, error) {
	client := fritzbox.NewClient(createHttpClient(config.IgnoreCertificates))

	endpoint, _ := url.Parse(config.FritzBoxUrl)
	client.BaseURL = endpoint

	err := client.Auth(config.FritzBoxUsername, config.FritzboxPassword)
	return client, err
}
