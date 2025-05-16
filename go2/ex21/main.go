package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Id       int64  `redis:"id"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
	Age      int64  `redis:"age"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.16:6379",
		Password: "goldMa$k46",
		DB:       0,
	})
	ctx := context.Background()
	ctx = context.WithValue(ctx, "rdb", rdb)
	users := []User{
		{Id: 1, Name: "BloomFilter", Email: "bf@email.address", Password: "111", Age: 21},
		{Id: 2, Name: "John Doe", Email: "jd@email.address", Password: "222", Age: 33},
		{Id: 3, Name: "Jane Doe", Email: "jane@email.address", Password: "333", Age: 27},
		{Id: 4, Name: "James Dean", Email: "james@email.address", Password: "444", Age: 38},
	}
	for i := 0; i < len(users); i++ {
		if err := CreateUser(ctx, users[i]); err != nil {
			panic(err)
		}
	}
}

func CreateUser(ctx context.Context, u User) error {
	irdb := ctx.Value("rdb")
	rdb, ok := irdb.(*redis.Client)
	if !ok {
		return fmt.Errorf("unable to get redis.Client from context")
	}
	cmds, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if err := pipe.HSet(ctx, fmt.Sprintf("users:%d", u.Id), u).Err(); err != nil {
			return err
		}
		if err := pipe.Set(ctx, fmt.Sprintf("users:email:%s", u.Email), u.Id, 0).Err(); err != nil {
			return err
		}
		if err := pipe.ZAdd(ctx, "users:age", redis.Z{Score: float64(u.Age), Member: u.Id}).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	for i := 0; i < len(cmds); i++ {
		fmt.Printf("%d-%d: %v\n", u.Id, i, cmds[i].Args())
	}
	return nil
}
