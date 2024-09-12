package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/integrations"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/log"
	"slices"
	"strconv"
	"strings"
)

const WhatsMyJobGif = tgbotapi.FileURL(integrations.NoOneHomeGifUrl)
const NoOneHomeGif = tgbotapi.FileURL(integrations.WhatsMyJobGifUrl)

type Integration struct {
	bot            *tgbotapi.BotAPI
	fritzBoxClient *fritzbox_requests.FritzBoxClientWithRefresh
}

func (i *Integration) createAPIClient() error {
	api, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err == nil {
		i.bot = api
	}
	return err
}

// IsEnabled returns true if the telegram feature is enabled.
// To be enabled the token and at least one allowed user must be configured
func IsEnabled() bool {
	return config.Get().TelegramBotToken != "" && len(config.Get().TelegramBotAllowedUsers) != 0
}

// New telegram client
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

func (i *Integration) reply(message *tgbotapi.Message, response string) {
	log.Print(log.CompTelegram, "Reply to user "+message.From.UserName)
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ParseMode = "HTML"
	msg.ReplyToMessageID = message.MessageID
	_, _ = i.bot.Send(msg)
}

func (i *Integration) send(msg tgbotapi.Chattable) {
	_, _ = i.bot.Send(msg)
}

func (i *Integration) sendGif(c *tgbotapi.Chat, url tgbotapi.FileURL) {
	giphy := tgbotapi.NewVideo(c.ID, url)
	_, _ = i.bot.Send(giphy)
}

func (i *Integration) start(update *tgbotapi.Update) {
	i.reply(update.Message, "Oh hey there! My only purpose is to tell you who is home.")
	i.sendGif(update.Message.Chat, WhatsMyJobGif)
}

func (i *Integration) whoIsOnline(u *tgbotapi.Update) {
	devices, err := fritzbox_requests.GetNetDevices(i.fritzBoxClient)
	if err != nil {
		i.reply(u.Message, "💥 Can not list online users, try again later")
		return
	}

	online := fritzbox_requests.MapToOnlineUsers(devices, false)

	if len(online) == 0 {
		i.reply(u.Message, "Oops! Looks like no one is home?")
		i.sendGif(u.Message.Chat, NoOneHomeGif)
		return
	}

	response := []string{"🏡 Currently home:\n"}
	for user, devices := range online {
		deviceCount := len(devices)
		devicesText := "device"
		if deviceCount > 1 {
			devicesText += "s"
		}
		response = append(response, "- <b>"+user+"</b> <i>online with "+strconv.Itoa(len(devices))+" "+devicesText+"</i>")
	}
	i.reply(u.Message, strings.Join(response, "\n"))
}

func (i *Integration) ListenForMessages() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := i.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			username := update.Message.From.UserName
			if !slices.Contains(config.Get().TelegramBotAllowedUsers, username) {
				i.reply(update.Message, "🤚 You are not allowed to use this bot!")
				continue
			}

			if !update.Message.IsCommand() {
				i.reply(update.Message, "💥 Please specify a valid command.")
				continue
			}

			switch update.Message.Command() {
			case "start":
				i.start(&update)
				break
			case "home", "online":
				i.whoIsOnline(&update)
				break
			default:
				i.reply(update.Message, "💥 Unknown command "+update.Message.Command())
				break
			}
		}
	}
}
