package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type User struct {
	Id       int64  `redis:"id"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
}

func main() {
	ctx := context.Background()
	user := &User{
		Id:       1,
		Name:     "BloomFilter",
		Email:    "bf@email.address",
		Password: "936a185caaa",
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.16:6379",
		Password: "goldMa$k46",
		DB:       0,
	})
	if err := rdb.HSet(ctx, "users:1", user).Err(); err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	output, err := rdb.HGetAll(ctx, "users:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
