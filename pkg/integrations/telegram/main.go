package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/config"
	"github.com/timo-reymann/fritzbox-based-presence/pkg/fritzbox_requests"
	"slices"
	"strconv"
	"strings"
)

type Integration struct {
	bot *tgbotapi.BotAPI
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
func New() (*Integration, error) {
	integration := Integration{}
	err := integration.createAPIClient()
	if err != nil {
		return nil, err
	}
	return &integration, nil
}

func (i *Integration) reply(message *tgbotapi.Message, response string) {
	println("[telegram-bot] Reply to user " + message.From.UserName)
	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ParseMode = "HTML"
	msg.ReplyToMessageID = message.MessageID
	_, _ = i.bot.Send(msg)
}

func (i *Integration) ListenForMessages(f *fritzbox_requests.FritzBoxClientWithRefresh) {
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
			case "home":
			case "online":
				i.reply(update.Message, getWhoIsHomeResponse(f))
				break
			default:
				i.reply(update.Message, "💥 Unknown command "+update.Message.Command())
				break
			}
		}
	}
}

func getWhoIsHomeResponse(f *fritzbox_requests.FritzBoxClientWithRefresh) string {
	devices, err := fritzbox_requests.GetNetDevices(f)
	if err != nil {
		return "💥 Can no list online users, try again later"
	}

	online := fritzbox_requests.MapToOnlineUsers(devices, false)
	response := []string{"🏡 Currently home:\n"}
	for user, devices := range online {
		deviceCount := len(devices)
		devicesText := "device"
		if deviceCount > 2 {
			devicesText += "s"
		}
		response = append(response, "- <b>"+user+"</b> <i>online with "+strconv.Itoa(len(devices))+" "+devicesText+"</i>")
	}
	return strings.Join(response, "\n")
}
