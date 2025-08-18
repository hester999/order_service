package redis

import (
	"context"
	"fmt"
	"order/internal/repo"
	"order/internal/repo/inmemmory/redis/order"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
	repo.OrderRepository
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client,
		OrderRepository: order.NewOrderRepository(client),
	}
}

func NewRedisClient(adr string, ctx context.Context) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr: adr,
	})
	if err := db.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}
	return db, nil
}
