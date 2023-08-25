package cmd

import (
	"fmt"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Run the CLI entrypoint
func Run() {
	// Read config
	err := config.Read()
	check(err)

	// Create client
	client, err := fritzbox_requests.CreateAuthenticatedFritzBoxClient(config.Get())
	check(err)

	// Get devices
	netDevices, err := fritzbox_requests.GetNetDevices(client)
	check(err)
	for _, device := range netDevices.Data.Active {
		fmt.Println(device)
	}
}
