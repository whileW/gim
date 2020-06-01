package utils

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool

var(
	Host = "47.104.27.123"
	Port = "6379"
)

func GetRedisPool() *redis.Pool {
	if pool == nil {
		pool = &redis.Pool{
			MaxIdle:     16,
			MaxActive:   0,
			IdleTimeout: 300,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", Host+":"+Port)
			},
		}
	}
	return pool
}
func Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply,err)
}
func Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply,err)
}