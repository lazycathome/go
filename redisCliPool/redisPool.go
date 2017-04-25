package redisCliPool

import (
	"time"

	"github.com/dlintw/goconf"
	"github.com/garyburd/redigo/redis"
)

var (
	// CliPool 连接池存储对象
	CliPool       *redis.Pool
	redisServer   string
	redisPassword string
	maxIdle       int
	maxActive     int
	timeout       int
	dbIndex       int
)

// initPool 初始化连接池
func initPool(server, password string, dbIndex int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 60 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			t := 60 * time.Second
			c, err := redis.Dial("tcp", server, redis.DialDatabase(dbIndex), redis.DialPassword(password), redis.DialReadTimeout(t), redis.DialWriteTimeout(t))
			if err != nil {
				return nil, err
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

//InitRedisPool 从配置文件中加载配置项初始化redis的连接池
func InitRedisPool(conf *goconf.ConfigFile) {
	redisServer, _ = conf.GetString("redisPool", "redisServer")
	redisPassword, _ = conf.GetString("redisPool", "redisPassword")
	maxIdle, _ = conf.GetInt("redisPool", "maxIdle")
	maxActive, _ = conf.GetInt("redisPool", "maxActive")
	dbIndex, _ = conf.GetInt("redisPool", "dbIndex")
	//idleTimeout, _ = conf.GetInt("redisPool", "idleTimeout")

	CliPool = initPool(redisServer, redisPassword, dbIndex)
	return
}

// String 对redis进行String命令封装
func String(reply interface{}, err1 error) (value string, err2 error) {
	value, err2 = redis.String(reply, err1)
	return
}
