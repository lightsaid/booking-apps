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
	// 1. 读取配置文件
	var config settings.AppConfig
	configs.NewConfig("config.yaml", &config, "./configs")

	// 2. 链接 PostgreSQL
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 3. 实例化 PostgreSQL 操作对象
	store := dbrepo.NewStore(db)

	// 4. 创建 Server
	server := NewServer(&config, store)

	// 5. 启动服务
	server.Start()
}
