package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type User struct {
	Id       int64  `redis:"id"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
	Age      int64  `redis:"age"`
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.16:6379",
		Password: "goldMa$k46",
		DB:       0,
	})
	users := []User{
		{Id: 1, Name: "BloomFilter", Email: "bf@email.address", Password: "111", Age: 21},
		{Id: 2, Name: "John Doe", Email: "jd@email.address", Password: "222", Age: 33},
		{Id: 3, Name: "Jane Doe", Email: "jane@email.address", Password: "333", Age: 27},
		{Id: 4, Name: "James Dean", Email: "james@email.address", Password: "444", Age: 38},
	}
	for i := 0; i < len(users); i++ {
		if err := CreateUser(rdb, ctx, users[i]); err != nil {
			panic(err)
		}
	}
	time.Sleep(100 * time.Millisecond)
	u1, err := FindUserByEmail(rdb, ctx, "jane@email.address")
	if err != nil {
		panic(err)
	}
	fmt.Printf("1st: %v\n", u1)
	u2, err := FindUserByAge(rdb, ctx, 20, 30)
	if err != nil {
		panic(err)
	}
	spew.Dump("2nd: %v\n", u2)
}

func CreateUser(rdb *redis.Client, ctx context.Context, u User) error {
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

func FindUser(rdb *redis.Client, ctx context.Context, id int64) (*User, error) {
	res, err := rdb.HGetAll(ctx, fmt.Sprintf("users:%d", id)).Result()
	if err != nil {
		return nil, err
	}
	user := new(User)
	user.Id, err = strconv.ParseInt(res["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	user.Age, err = strconv.ParseInt(res["age"], 10, 64)
	if err != nil {
		return nil, err
	}
	user.Name = res["name"]
	user.Email = res["email"]
	user.Password = res["password"]
	return user, nil
}

func FindUserByEmail(rdb *redis.Client, ctx context.Context, email string) (*User, error) {
	res, err := rdb.Get(ctx, "users:email:"+email).Int64()
	if err != nil {
		return nil, err
	}
	user, err := FindUser(rdb, ctx, res)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByAge(rdb *redis.Client, ctx context.Context, startAge int64, endAge int64) ([]*User, error) {
	res, err := rdb.ZRangeByScore(ctx, "users:age", &redis.ZRangeBy{Min: strconv.FormatInt(startAge, 10), Max: strconv.FormatInt(endAge, 10)}).Result()
	if err != nil {
		return nil, err
	}
	users := make([]*User, 0)
	for i := 0; i < len(res); i++ {
		id, err := strconv.ParseInt(res[i], 10, 64)
		if err != nil {
			return nil, err
		}
		user, err := FindUser(rdb, ctx, id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
