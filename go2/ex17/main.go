package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.16:6379",
		Password: "goldMa$k46",
		DB:       0,
	})
	if err := rdb.Set(ctx, "a1", "Hello", time.Hour).Err(); err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "a1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
