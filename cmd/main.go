package cmd

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/buildinfo"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/integrations/telegram"
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

	// Create telegram client
	if telegram.IsEnabled() {
		integration, err := telegram.New()
		if err != nil {
			println("[telegram-bot] Failed to start bot feature: " + err.Error())
		} else {
			println("[telegram-bot] Listening for messages ...")
			go integration.ListenForMessages(client)
		}
	}

	println("[server] Spinning up ...")
	err = server.Start(config.Get(), client)
	if err != nil {
		println("[server] Failed to startup server: " + err.Error())
	}
}
