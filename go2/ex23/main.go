package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "192.168.0.16:6379", Password: "goldMa$k46"})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	args := os.Args[1:]
	if len(args) > 0 {
		if duration, err := time.ParseDuration(args[0]); err == nil {
			go func() {
				time.Sleep(duration)
				cancel()
			}()
		}
	}
	if err := rdb.Do(ctx, "debug", "sleep", "5").Err(); err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
