package main

import (
	"database/sql"
	"log"
	"toolkit/configs"

	_ "github.com/lib/pq"

	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/settings"
)

func main() {
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := dbrepo.NewStore(db)

	server := NewServer(&config, store)

	server.Start()
}
