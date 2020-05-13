package config

import (
	logConfig "cy/lib"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"time"
)

var LocalMysql *gorm.DB
var ApiConfig *ApiCfg

func Load() {

	if err := initConfig(); err != nil {

	}

	ApiConfig = initApiCfg()
	LocalMysql = initMysql()

	//接收请求日志设置
	_ = initLog()

	//设置时区
	_, _ = time.LoadLocation("Asia/Shanghai")
}

func initConfig() error {

	viper.AddConfigPath("./")
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func initApiCfg() *ApiCfg {
	cfg := &ApiCfg{
		HostEnv: viper.GetString("env"),
	}

	return cfg
}

func initMysql() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.local.user"),
		viper.GetString("mysql.local.pass"),
		viper.GetString("mysql.local.host"),
		viper.GetString("mysql.local.port"),
		viper.GetString("mysql.local.name"),
	)

	db, err := gorm.Open("mysql", dns)

	if err != nil {
		log.Fatal("Connecting mysql error...")
	}

	return db
}

func initLog() error {
	//配置日志
	logConf := logConfig.LogConfig{
		Level: viper.GetString("log.log_level"),
		FW: logConfig.ConfFileWriter{
			On:              viper.GetBool("log.file_writer.log_on"),
			LogPath:         viper.GetString("log.file_writer.log_path"),
			RotateLogPath:   viper.GetString("log.file_writer.rotate_log_path"),
			WfLogPath:       viper.GetString("log.file_writer.wf_log_path"),
			RotateWfLogPath: viper.GetString("log.file_writer.rotate_wf_log_path"),
		},
		CW: logConfig.ConfConsoleWriter{
			On:    viper.GetBool("log.console_writer.log_on"),
			Color: viper.GetBool("log.console_writer.color"),
		},
	}

	if err := logConfig.SetupDefaultLogWithConf(logConf); err != nil {
		panic(err)
	}

	logConfig.SetLayout("2006-01-02T15:04:05.000")

	return nil
}
