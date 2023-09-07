package config

import (
	"github.com/kelseyhightower/envconfig"
)

// GuestsUsername is the collective name to be displayed for devices that are currently online
const GuestsUsername = "Guests"

// AppConfig represents the configuration for the entire application
type AppConfig struct {
	// FritzBoxUrl is the base URL for the Fritz!Box Web UI
	FritzBoxUrl string `required:"true" split_words:"true"`
	// FritzBoxUsername to use for logging in
	FritzBoxUsername string `required:"true" split_words:"true"`
	// FritzBoxPassword to use for logging in
	FritzBoxPassword string `required:"true" split_words:"true"`
	// IgnoreCertificates specifies to ignore SSL certificate validation for the requests to Fritz!Box
	IgnoreCertificates bool `required:"false" default:"false" split_words:"true"`
	// DeviceNameMapping is the list of device names to map to given usernames
	DeviceNameMapping DeviceNameMappingDecoder `required:"true" split_words:"true"`
	// ServerPort is the HTTP port of the server
	ServerPort int `required:"false" default:"8090" split-word:"true"`
	// ShowGuests enables the visibility of guest devices, bundled under a pseudo username
	ShowGuests bool `required:"false" default:"true" split_words:"true"`
	// AuthMiddlewareOrder is a ordered list, which is followed to authenticate users
	AuthMiddlewareOrder []string `required:"false" default:"ip_range,www_authenticate"  split_words:"true"`
	// AuthIpRange is the local IP range to allow direct access to the page
	AuthIpRange IpRangeDecoder `required:"false" default:"192.168.178.0/24"  split_words:"true"`
	// AuthPassword is the shared password for WWW-Authenticate to get access
	AuthPassword string `required:"false" default:"changeme" split_words:"true"`
	// IndexTemplatePath is the location of a custom UI template to use
	IndexTemplatePath string `required:"false" split_words:"true"`
	// TelegramBotToken is the token to use for doing telegram requests. If it is set telegram messages are read
	TelegramBotToken string `required:"false" split_words:"true"`
	// TelegramBotAllowedUsers is the list of usernames allowed to talk to the bot
	TelegramBotAllowedUsers []string `required:"false" split_words:"true"`
}

var config AppConfig

// Read from env variables
func Read() error {
	c := AppConfig{}
	err := envconfig.Process("fb_presence_", &c)
	config = c
	return err
}

// PrintUsage displays the help for the env vars
func PrintUsage() {
	_ = envconfig.Usage("fb_presence_", &AppConfig{})
}

// Get current app configuration, make sure Read has been called before.
func Get() *AppConfig {
	return &config
}
