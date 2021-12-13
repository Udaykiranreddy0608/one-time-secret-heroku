package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Test() string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "ping", time.Now(), 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "ping").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ping", val)
	return val
}
