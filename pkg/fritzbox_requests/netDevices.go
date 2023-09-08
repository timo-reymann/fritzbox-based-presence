package fritzbox_requests

import (
	"net/http"
	"net/url"
)

// NetDevicesResponse represents the Fritz!Box data request to get all network devices known
type NetDevicesResponse struct {
	Data struct {
		Active []struct {
			Mac   string `json:"mac"`
			Name  string `json:"name"`
			State struct {
				Class string `json:"class"`
			} `json:"state"`
		} `json:"active"`
	} `json:"data"`
}

// GetNetDevices loads all known devices from fritzbox using the specified client
func GetNetDevices(c *FritzBoxClientWithRefresh) (response *NetDevicesResponse, err error) {
	response = &NetDevicesResponse{}

	err = DoWithRetry(c, func() *http.Request {
		v := url.Values{}
		v.Set("page", "netDev")
		v.Set("xhrId", "cleanup")
		req, _ := c.NewRequest("POST", "/data.lua", v)
		return req
	}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
