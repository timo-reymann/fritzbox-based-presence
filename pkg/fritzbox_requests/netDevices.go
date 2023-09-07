package fritzbox_requests

import (
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
	v := url.Values{}
	v.Set("page", "netDev")
	v.Set("xhrId", "cleanup")

	req, err := c.NewRequest("POST", "/data.lua", v)
	if err != nil {
		return nil, err
	}

	response = &NetDevicesResponse{}

	// retry to get active devices as it seems the fritz!box sometimes
	// returns an empty list for active devices, especially after refreshing
	// the session.
	i := 0
	for i < 2 && response.Data.Active == nil {
		err = DoWithRetry(c, req, response)
		i++
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}
