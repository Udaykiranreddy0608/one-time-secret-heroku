package goredis

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func InitPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func Get(key string) (string, string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
		return "", "", err
	}

	return key, s, nil
}

func Set(key string, val string) (string, string, error) {
	// get conn and put back when exit from method
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return "", "", err
	}

	return key, val, nil
}
