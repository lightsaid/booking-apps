package main

import (
	"database/sql"
	"log"
	"toolkit/configs"

	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/settings"
)

// 此程序是为了快速生成一些样板数据

var store dbrepo.Store

func main() {
	// 1. 读取配置文件
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	// 2. 链接 PostgreSQL
	db := connectPostgreSQL(&config)
	defer db.Close()

	// 3. 实例化 PostgreSQL 操作对象
	store = dbrepo.NewStore(db)
}

func connectPostgreSQL(config *settings.AppConfig) *sql.DB {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
