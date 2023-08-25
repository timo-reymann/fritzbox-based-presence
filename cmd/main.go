package cmd

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/buildinfo"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Run the CLI entrypoint
func Run() {
	println("Getting build information ...")
	buildinfo.PrintVersionInfo()
	// Read config
	err := config.Read()
	if err != nil {
		println("Error loading configuration: " + err.Error())
		config.PrintUsage()
		return
	}

	// Create client
	println("Creating Fritz!Box session ...")
	client, err := fritzbox_requests.CreateAuthenticatedFritzBoxClient(config.Get())
	if err != nil {
		println("Failed to authenticate with Fritz!Box: " + err.Error())
		return
	}

	println("Spinning up server ...")
	err = server.Start(config.Get(), client)
	if err != nil {
		println("Failed to startup server: " + err.Error())
	}
}
