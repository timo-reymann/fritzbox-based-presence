package fritzbox_requests

import (
	"crypto/tls"
	"errors"
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
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

// DoWithRetry executes the given do request for the fritzbox client
func DoWithRetry[T interface{}](c *FritzBoxClientWithRefresh, reqFactory func() *http.Request, res *T) error {
	t := new(T)
	_, err := c.Do(reqFactory(), t)

	// Retry by authenticating again
	if errors.Is(err, fritzbox.ErrExpiredSess) {
		log.Print(log.CompFritzbox, "Refresh session")
		err = c.refreshSession()
		if err != nil {
			return err
		}
		return DoWithRetry(c, reqFactory, res)
	}

	if err == nil {
		*res = *t
	}

	return err
}

// Do executes the given request. If it leads to a expired session error
// it updates the session
func (c *FritzBoxClientWithRefresh) Do(req *http.Request, v interface{}) (*http.Response, error) {
	log.Print(log.CompFritzbox, req.Method+" "+req.URL.Path)
	return c.fritzBoxClient.Do(req, v)
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
