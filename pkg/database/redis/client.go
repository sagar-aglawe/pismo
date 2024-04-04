package redis

import (
	"context"
	"fmt"

	redisPkg "github.com/redis/go-redis/v9"
)

type Config struct {
	Address string
}

func New(config *Config) {
	rdb := redisPkg.NewClient(&redisPkg.Options{
		Addr: config.Address,
	})

	err := rdb.Set(context.Background(), "key", "sagar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
