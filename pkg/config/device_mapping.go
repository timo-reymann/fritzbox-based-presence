package config

import (
	"errors"
	"strings"
)

type DeviceNameMappingDecoder map[string][]string

func (umd *DeviceNameMappingDecoder) Decode(value string) error {
	if len(value) == 0 {
		return errors.New("must be longer than 1 character")
	}

	parsed := map[string][]string{}

	devicesPerUser := strings.Split(value, "|")
	for _, devicePerUser := range devicesPerUser {
		devicesAndUser := strings.Split(devicePerUser, "=")
		if len(devicesAndUser) != 2 {
			return errors.New("expected device mapping to in form <user>=<device1>[,<device2>]")
		}

		user := devicesAndUser[0]
		devices := strings.Split(devicesAndUser[1], ",")

		if len(devices) == 0 || devices[0] == "" {
			return errors.New("every user needs to have at least one device assigned")
		}

		parsed[user] = make([]string, len(devices))
		for idx, device := range devices {
			parsed[user][idx] = device
		}
	}

	*umd = parsed
	return nil
}

func IsDeviceFor(device string) string {
	for user, devices := range config.DeviceNameMapping {
		for _, deviceFromConfig := range devices {
			if deviceFromConfig == device {
				return user
			}
		}
	}

	return ""
}
