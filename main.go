package main

import (
	"context"
	"github.com/iralance/coredemo/framework"
	"github.com/iralance/coredemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())
	registerRouter(core)
	server := http.Server{
		Addr:    ":8080",
		Handler: core,
	}
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT ctrl+c kill+\
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
