package cmd

import (
	"github.com/timo-reymann/fritzbox-based-presence/pkg/buildinfo"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/integrations/discord"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/integrations/telegram"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/server"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Run the CLI entrypoint
func Run() {
	log.Print(log.CompCli, "Getting build information ...")
	buildinfo.PrintVersionInfo()
	// Read config
	err := config.Read()
	if err != nil {
		log.Print(log.CompCli, "Error loading configuration: "+err.Error())
		config.PrintUsage()
		return
	}

	// Create client
	log.Print(log.CompFritzbox, "Creating Fritz!Box session ...")
	client, err := fritzbox_requests.CreateAuthenticatedFritzBoxClient(config.Get())
	if err != nil {
		log.Print(log.CompFritzbox, "Failed to authenticate with Fritz!Box: "+err.Error())
		return
	}

	// Create telegram client
	if telegram.IsEnabled() {
		integration, err := telegram.New(client)
		if err != nil {
			log.Print(log.CompTelegram, "Failed to start telegram bot feature: "+err.Error())
		} else {
			log.Print(log.CompTelegram, "Listening for telegram messages ...")
			go integration.ListenForMessages()
		}
	}

	if discord.IsEnabled() {
		_, err := discord.New(client)
		if err != nil {
			log.Print(log.CompDiscord, "Failed to start discord bot feature: "+err.Error())
		} else {
			log.Print(log.CompDiscord, "Listening for discord messages ...")
		}
	}

	log.Print(log.CompServer, "Spinning up ...")
	err = server.Start(config.Get(), client)
	if err != nil {
		log.Print(log.CompServer, "Failed to startup server: "+err.Error())
	}
}
