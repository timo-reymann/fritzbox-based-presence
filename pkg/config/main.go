package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	FritzBoxUrl        string `required:"true" envconfig:"fb_presence__fritzbox_url"`
	FritzBoxUsername   string `required:"true" envconfig:"fb_presence__fritzbox_username"`
	FritzboxPassword   string `required:"true" envconfig:"fb_presence__fritzbox_password"`
	IgnoreCertificates bool   `required:"false" default:"false" envconfig:"fb_presence__ignore_certificates"`
}

var config AppConfig

// Read from env variables
func Read() error {
	c := AppConfig{}
	err := envconfig.Process("", &c)
	config = c
	return err
}

// Get current app configuration, make sure Read has been called before.
func Get() *AppConfig {
	return &config
}
