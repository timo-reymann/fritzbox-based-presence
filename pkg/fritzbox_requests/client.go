package fritzbox_requests

import (
	"crypto/tls"
	"errors"
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

type FritzBoxClientWithRefresh struct {
	fritzBoxClient *fritzbox.Client
	endpoint       *url.URL
	username       string
	password       string
	httpClient     *http.Client
}

func NewFritzBoxClientWithRefresh(c *http.Client, endpoint *url.URL) FritzBoxClientWithRefresh {
	refreshClient := FritzBoxClientWithRefresh{
		httpClient: c,
		endpoint:   endpoint,
	}
	refreshClient.createClient()
	return refreshClient
}

func (c *FritzBoxClientWithRefresh) createClient() {
	c.fritzBoxClient = fritzbox.NewClient(c.httpClient)
	c.fritzBoxClient.BaseURL = c.endpoint
}

// Do executes the given request. If it leads to a expired session error
// it updates the session
func (c *FritzBoxClientWithRefresh) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.fritzBoxClient.Do(req, v)

	// Retry by authenticating again
	if errors.Is(err, fritzbox.ErrExpiredSess) {
		err := c.refreshSession()
		if err != nil {
			return nil, err
		}

		return c.fritzBoxClient.Do(req, v)
	}

	return res, err
}

// Auth sends a auth request and returns an error, if any. Session is stored
// in client in order to perform requests with authentification.
func (c *FritzBoxClientWithRefresh) Auth(username string, password string) error {
	c.username = username
	c.password = password
	return c.fritzBoxClient.Auth(username, password)
}

func (c *FritzBoxClientWithRefresh) refreshSession() error {
	c.createClient()
	return c.fritzBoxClient.Auth(c.username, c.password)
}

// NewRequest creates an API request. A relative URL can be provided
// in urlStr in which case it is resolved relative to the BaseURL of
// the Client. Relative URLs should always be specified without a
// preceding slash. If specified, the value pointed to by data is Query
// encoded and included as the request body in order to perform form requests.
func (c *FritzBoxClientWithRefresh) NewRequest(method, urlStr string, data url.Values) (*http.Request, error) {
	return c.fritzBoxClient.NewRequest(method, urlStr, data)
}

// CreateAuthenticatedFritzBoxClient generates an fritzbox client based on given configuration.
// After parsing the url a login is attempted, if that fails an error is returned
func CreateAuthenticatedFritzBoxClient(config *config.AppConfig) (*FritzBoxClientWithRefresh, error) {
	endpoint, _ := url.Parse(config.FritzBoxUrl)
	client := NewFritzBoxClientWithRefresh(createHttpClient(config.IgnoreCertificates), endpoint)

	err := client.Auth(config.FritzBoxUsername, config.FritzBoxPassword)
	return &client, err
}
