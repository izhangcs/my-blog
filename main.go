package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
	"zhangcs/blog/config"
	"zhangcs/blog/model"
	v "zhangcs/blog/pkg/version"
	"zhangcs/blog/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "配置文件路径")
	version = pflag.BoolP("version", "v", false, "显示版本信息")
)

func main() {
	// 初始化配置
	pflag.Parse()

	// 显示版本信息
	if *version {
		showVersion()
		return
	}

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

	// 设置静态资源
	g.LoadHTMLGlob("views/**/*")
	g.Static("/static", "static")
	// 开启 session
	store := cookie.NewStore([]byte("scret"))
	g.Use(sessions.Sessions("blog", store))

	middlwares := []gin.HandlerFunc{}

	router.Load(g, middlwares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
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

func showVersion() {
	v := v.Get()
	marshalled, err := json.MarshalIndent(&v, "", "  ")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(marshalled))
}
