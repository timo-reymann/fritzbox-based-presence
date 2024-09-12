package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/integrations"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"

	"strconv"
	"strings"
)

const processingEmoji = "⌛"

type Integration struct {
	session        *discordgo.Session
	fritzBoxClient *fritzbox_requests.FritzBoxClientWithRefresh
}

// IsEnabled returns true if the discord feature is enabled.
// To be enabled the token and at least one allowed user must be configured
func IsEnabled() bool {
	return config.Get().DiscordBotToken != ""
}

func (i *Integration) createAPIClient() error {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", config.Get().DiscordBotToken))
	if err != nil {
		return err
	}

	i.session = session
	i.session.AddHandler(i.listenForMessage)

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Print(log.CompDiscord, "Logged in as "+r.User.String())
	})

	err = i.session.Open()
	if err != nil {
		return err
	}
	return nil
}

func (i *Integration) respond(m *discordgo.MessageCreate, content string) {
	_ = i.session.MessageReactionRemove(m.ChannelID, m.Message.ID, processingEmoji, i.session.State.User.ID)
	_, _ = i.session.ChannelMessageSend(m.ChannelID, content)
}

func (i *Integration) isAllowedUser(u *discordgo.User) bool {
	for _, user := range config.Get().DiscordBotAllowedUsers {
		if u.Username == user || u.ID == user {
			return true
		}
	}

	return false
}

func (i *Integration) listenForMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	_ = i.session.MessageReactionAdd(m.ChannelID, m.Message.ID, processingEmoji)

	if !i.isAllowedUser(m.Author) {
		i.respond(m, ":octagonal_sign: You are not allowed to use this bot.")
		return
	}

	if m.Content != "/home" && m.Content != "/online" {
		i.respond(m, "💥 Unknown command. Try /home or /online")
		return
	}

	i.listUsers(m)
}

func (i *Integration) listUsers(m *discordgo.MessageCreate) {
	devices, err := fritzbox_requests.GetNetDevices(i.fritzBoxClient)
	if err != nil {
		i.respond(m, "💥 Can not list online users, try again later")
		return
	}

	online := fritzbox_requests.MapToOnlineUsers(devices, false)

	if len(online) == 0 {
		i.respond(m, "Oops! Looks like no one is home?")
		i.respond(m, integrations.NoOneHomeGifUrl)
		return
	}

	response := []string{"🏡 Currently home:"}
	for user, devices := range online {
		deviceCount := len(devices)
		devicesText := "device"
		if deviceCount > 1 {
			devicesText += "s"
		}
		response = append(response, "- **"+user+"** *online with "+strconv.Itoa(len(devices))+" "+devicesText+"*")
	}
	i.respond(m, strings.Join(response, "\n"))
}

func New(f *fritzbox_requests.FritzBoxClientWithRefresh) (*Integration, error) {
	integration := Integration{
		fritzBoxClient: f,
	}
	err := integration.createAPIClient()
	if err != nil {
		return nil, err
	}
	return &integration, nil
}
