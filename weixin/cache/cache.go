//Package cache 缓存管理
package cache

import (
	"lazycathome/redisCliPool"

	"github.com/garyburd/redigo/redis"
)

//SetGobalToken 设置全局token
func SetGobalToken(key, value string, expireTime int) bool {
	conn := redisCliPool.CliPool.Get()
	r, _ := redis.String(conn.Do("SETEX", key, expireTime, value))
	result := false
	if "OK" == r {
		result = true
	}
	return result
}

//SetAuthToken 设置授权token
func SetAuthToken(key, value string, expireTime int) bool {
	conn := redisCliPool.CliPool.Get()
	r, _ := redis.String(conn.Do("SETEX", key, expireTime, value))
	result := false
	if "OK" == r {
		result = true
	}
	return result
}
