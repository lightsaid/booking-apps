package main

import (
	"fmt"
	"toolkit/configs"

	"github.com/lightsaid/booking-sys/pkg/settings"
)

func main() {
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	fmt.Println(
		config.DBDriver, "\n",
		config.DBSource, "\n",
		config.Port, "\n",
		config.Secret)
}
