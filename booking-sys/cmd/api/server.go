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

	"github.com/gin-gonic/gin"
	"github.com/lightsaid/booking-sys/pkg/settings"
)

type Server struct {
	config *settings.AppConfig
	router *gin.Engine
}

func NewServer(config *settings.AppConfig) *Server {
	server := &Server{
		config: config,
	}
	server.initRouter()

	return server
}

func (server *Server) Start() {
	s := http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", server.config.Server.Port),
		Handler:        server.router,
		IdleTimeout:    time.Minute,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 4 << 20, // 4M
	}

	// quitErr := make(chan error)

	// go func() {
	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 	// 阻塞，等待信息号
	// 	sInfo := <-quit
	// 	log.Println("Stop server signal: ", sInfo.String())

	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()

	// 	log.Println("Stopping server...")

	// 	err := s.Shutdown(ctx)
	// 	if err != nil {
	// 		fmt.Println("ssssss")
	// 		quitErr <- err
	// 	}
	// 	fmt.Println("bbbbbbb")
	// 	quitErr <- nil
	// }()

	// log.Println("Starting server on ", s.Addr)
	// err := s.ListenAndServe()
	// if err != nil {
	// 	fmt.Println("ccccc")
	// 	return err
	// }

	// // 阻塞，等待 quitErr channel 错误信号
	// err = <-quitErr
	// fmt.Println("接受quitErr: ", err)
	// if err != nil {
	// 	return err
	// }

	// // 成功关机
	// log.Println("Stopped server.")

	// return nil

	go func() {
		log.Println("Starting server on ", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			log.Println("ListenAndServe: ", err)
		}
	}()

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
