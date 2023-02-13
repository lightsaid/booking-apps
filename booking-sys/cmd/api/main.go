package main

import (
	"toolkit/configs"

	"github.com/lightsaid/booking-sys/pkg/settings"
)

func main() {
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	server := NewServer(&config)

	server.Start()
}
