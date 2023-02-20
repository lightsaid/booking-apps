package main

import (
	"database/sql"
	"log"
	"toolkit/configs"

	_ "github.com/lib/pq"

	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	_ "github.com/lightsaid/booking-sys/docs"
	"github.com/lightsaid/booking-sys/pkg/settings"
)

// @title Booking System API
// @version 1.0
// @description 这是电影院购票系统的后端服务API
// @termsOfService https://github.com/lightsaid/booking-apps

// @host localhost:5000
// @BasePath /v1/api
func main() {
	// 1. 读取配置文件
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	// 2. 链接 PostgreSQL
	db := connectPostgreSQL(&config)
	defer db.Close()

	// 3. 实例化 PostgreSQL 操作对象
	store := dbrepo.NewStore(db)

	// 4. 创建 Server
	server := NewServer(&config, store)

	// 5. 启动服务
	server.Start()
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
