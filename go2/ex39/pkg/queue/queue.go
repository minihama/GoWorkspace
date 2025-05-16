package queue

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type (
	Order struct {
		UserId   int       `json:"user_id"`
		OrderId  int       `json:"order_id"`
		Products []Product `json:"products"`
	}
	Product struct {
		ProductId int     `json:"product_id"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	}
)

type RedisQueue struct {
	client *redis.Client
}

func New() *RedisQueue {
	rdb := redis.NewClient(&redis.Options{Addr: "192.168.0.16:6379", Password: "goldMa$k46"})
	return &RedisQueue{client: rdb}
}

func (rq *RedisQueue) Enqueue(val *Order) error {
	order, err := json.Marshal(val)
	if err != nil {
		return err
	}
	cmd := rq.client.RPush(context.Background(), "order", string(order))
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (rq *RedisQueue) Dequeue() (*Order, error) {
	cmd := rq.client.LPop(context.Background(), "order")
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	order := new(Order)
	if err := json.Unmarshal([]byte(cmd.Val()), order); err != nil {
		return nil, err
	}
	return order, nil
}
