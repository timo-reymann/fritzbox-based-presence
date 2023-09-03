package config

import (
	"github.com/kelseyhightower/envconfig"
)

// GuestsUsername is the collective name to be displayed for devices that are currently online
const GuestsUsername = "Guests"

// AppConfig represents the configuration for the entire application
type AppConfig struct {
	FritzBoxUrl         string                   `required:"true" split_words:"true"`
	FritzBoxUsername    string                   `required:"true" split_words:"true"`
	FritzBoxPassword    string                   `required:"true" split_words:"true"`
	IgnoreCertificates  bool                     `required:"false" default:"false" split_words:"true"`
	DeviceNameMapping   DeviceNameMappingDecoder `required:"true" split_words:"true"`
	ServerPort          int                      `required:"false" default:"8090"`
	ShowGuests          bool                     `required:"false" default:"true" split_words:"true"`
	AuthIpRange         IpRangeDecoder           `required:"false" default:"192.168.178.0/24"  split_words:"true"`
	AuthMiddlewareOrder []string                 `required:"false" default:"ip_range"  split_words:"true"`
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
