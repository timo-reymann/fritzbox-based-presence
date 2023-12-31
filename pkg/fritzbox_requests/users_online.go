package fritzbox_requests

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
)

// UsersOnlineMapping is indexed by the human-readable username
// and a list of all devices currently connected for them
type UsersOnlineMapping = map[string][]string

func CreateUsersOnlineMapping() UsersOnlineMapping {
	return UsersOnlineMapping{}
}

// MapToOnlineUsers takes a given net devices response and transforms it into the user mapping
func MapToOnlineUsers(netDevicesRes *NetDevicesResponse, includeOffline bool) UsersOnlineMapping {
	usersOnline := CreateUsersOnlineMapping()

	if includeOffline {
		if config.Get().ShowGuests {
			usersOnline[config.GuestsUsername] = []string{}
		}

		for username := range config.Get().DeviceNameMapping {
			usersOnline[username] = []string{}
		}
	}

	for _, device := range netDevicesRes.Data.Active {
		user := config.IsDeviceFor(device.Name)

		if config.Get().ShowGuests && device.State.Class == "globe_online_guest" {
			user = config.GuestsUsername
		}

		if user == "" {
			continue
		}

		_, ok := usersOnline[user]
		if !ok {
			usersOnline[user] = []string{}
		}
		usersOnline[user] = append(usersOnline[user], device.Name)
	}
	return usersOnline
}
