package main

import (
	"cy/config"
	"cy/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Load()

	router.HttpServerRun()

	/**
	@date 2020-04-25
	@desc 终止服务
	*/
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
