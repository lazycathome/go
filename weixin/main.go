package main

import (
	"flag"
	"fmt"
	"lazycathome/redisCliPool"
	"lazycathome/weixin/net"
	"os"

	"github.com/dlintw/goconf"
)

// wxConf 微信配置
var wxConf WXConf

func main() {

	confFile := loadConf()
	//加载redis配置文件，且初始化redis连接池
	initRedis(confFile)
	//初始化微信配置
	initWxConf(confFile)
	fmt.Println(wxConf.AppId)
	_, err := net.Get("http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	var a string = "c"
	if a != "" && a != "a" {
		fmt.Println("test")
	} else {
		fmt.Println("test2")
	}
	//fmt.Println(string(result))

}

func loadConf() *goconf.ConfigFile {
	confFile := flag.String("config", "conf.ini", "set redis config file.")
	flag.Parse()

	var err error

	conf, err := goconf.ReadConfigFile(*confFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", confFile, err)
		os.Exit(1)
	}
	return conf
}

func initWxConf(conf *goconf.ConfigFile) {
	appId, _ := conf.GetString("weixin", "appId")
	appSecret, _ := conf.GetString("weixin", "appSecret")
	wxConf.AppId = appId
	wxConf.AppSecret = appSecret
}

func initRedis(conf *goconf.ConfigFile) {

	redisCliPool.InitRedisPool(conf)
}
