package repository

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func NewRedisClient(addr string) (*redis.Client, error) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: "",
        DB:       0,
    })

    ctx := context.Background()
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        return nil, err
    }

    return rdb, nil
}