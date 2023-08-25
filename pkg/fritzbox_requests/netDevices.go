package fritzbox_requests

import (
	"github.com/philippfranke/go-fritzbox/fritzbox"
	"net/url"
)

// NetDevicesResponse represents the Fritz!Box data request to get all network devices known
type NetDevicesResponse struct {
	Data struct {
		Active []struct {
			Mac   string `json:"mac"`
			Name  string `json:"name"`
			State string `json:"state"`
		} `json:"active"`
	} `json:"data"`
}

// GetNetDevices loads all known devices from fritzbox using the specified client
func GetNetDevices(c *fritzbox.Client) (response *NetDevicesResponse, err error) {
	v := url.Values{}
	v.Set("page", "netDev")
	v.Set("lang", "de")
	v.Set("xhrId", "cleanup")

	req, err := c.NewRequest("POST", "/data.lua", v)
	if err != nil {
		return nil, err
	}

	response = &NetDevicesResponse{}
	_, err = c.Do(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
