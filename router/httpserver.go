package router

import (
	"context"
	"cy/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	if config.ApiConfig.HostEnv == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           viper.GetString("http.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(viper.GetInt("http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(viper.GetInt("http.max_header_bytes")),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", viper.GetString("http.addr"))
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", viper.GetString("http.addr"), err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
