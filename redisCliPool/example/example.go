package main

import (
	"flag"
	"fmt"
	"lazycathome/redisCliPool"
	"os"
	"strconv"

	"github.com/dlintw/goconf"
	"github.com/garyburd/redigo/redis"
)

var channel chan argValue

var sChannel chan string

func main() {
	var err error
	confFile := flag.String("config", "conf.ini", "set redis config file.")
	flag.Parse()

	conf, err := goconf.ReadConfigFile(*confFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "LoadConfiguration: Error: Could not open %q for reading: %s\n", confFile, err)
		os.Exit(1)
	}

	redisCliPool.InitRedisPool(conf)
	//defer redisCliPool.CliPool.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	conn := redisCliPool.CliPool.Get()
	defer conn.Close()

	r2, _ := redis.String(conn.Do("GET", "name"))
	fmt.Println(r2)

	channel = make(chan argValue)

	sChannel = make(chan string)

	for i := 0; i < 10; i++ {
		go getWorker(i)
	}

	for i := 0; i < 10; i++ {
		go setWorker(i)
	}

	for index := 0; index < 100; index++ {
		s := strconv.Itoa(index)
		kv := argValue{"name" + s, "red" + s}
		channel <- kv
		//go set("name"+s, "red"+s)
	}

	for index := 0; index < 100; index++ {
		s := strconv.Itoa(index)
		//go get("name" + s)
		sChannel <- "name" + s
	}

	select {}
}

type argValue struct {
	name  string
	value string
}

func getWorker(i int) {
	for {
		get(<-sChannel, i)
	}
}

func setWorker(i int) {
	for {
		set(<-channel, i)
	}
}

func get(key string, i int) string {
	var err error
	var v string
	conn := redisCliPool.CliPool.Get()
	defer conn.Close()

	tv, err := conn.Do("GET", key)
	v, err = redisCliPool.String(tv, err)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println("channel", i, ", key:", v)
	return v
}

func set(kv argValue, i int) {
	conn := redisCliPool.CliPool.Get()
	defer conn.Close()
	r, err := conn.Do("SET", kv.name, kv.value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("key: %s, value:%s,  r:%s", kv.name, kv.value, r)
	fmt.Println()
}
