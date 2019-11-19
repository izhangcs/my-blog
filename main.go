package main

import (
	"errors"
	"net/http"
	"time"
	"zhangcs/blog/config"
	"zhangcs/blog/model"
	"zhangcs/blog/router"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "配置文件路径")
)

func main() {
	// 初始化配置
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	//创建gin
	g := gin.New()
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	middlwares := []gin.HandlerFunc{}

	router.Load(g, middlwares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
