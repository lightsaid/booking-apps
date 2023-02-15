package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toolkit/logz"

	"github.com/gin-gonic/gin"
	dbrepo "github.com/lightsaid/booking-sys/dbrepo/postgres"
	"github.com/lightsaid/booking-sys/pkg/settings"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	config *settings.AppConfig
	router *gin.Engine
	store  dbrepo.Store
}

func NewServer(config *settings.AppConfig, store dbrepo.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
	}
	server.initRouter()

	return server
}

func (server *Server) Start() {
	// 1. 设置日志
	var logLevel zapcore.Level
	if server.config.Server.RunMode == "release" {
		logLevel = zapcore.InfoLevel
	} else {
		logLevel = zapcore.DebugLevel
	}
	logz.SettingGlobalLogger(server.config.LogOutput, logLevel)

	// 2. 设置 validator 引擎
	err := setupValidatorEngine()
	if err != nil {
		log.Fatal(err)
	}

	// 3. 创建 HTTP Server
	s := http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", server.config.Server.Port),
		Handler:        server.router,
		IdleTimeout:    time.Minute,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 4 << 20, // 4M
	}

	// 启动服务监听
	go func() {
		log.Println("Starting server on ", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			log.Println("ListenAndServe: ", err)
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//  阻塞，等待 os.Signal 信号
	<-quit
	log.Println("Stopping server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown server error: ", err)
	}

	log.Println("Stopped server.")
}
