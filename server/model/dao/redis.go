package dao

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func InitPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	RedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
