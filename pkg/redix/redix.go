package redix

import (
	"context"
	"errors"
	"fmt"

	"github.com/diazharizky/go-app-core/config"
	"github.com/redis/go-redis/v9"
)

type Redix struct {
	client *redis.Client
}

func init() {
	config.Global.SetDefault("redis.host", "localhost")
	config.Global.SetDefault("redis.port", 6379)
}

func New() (*Redix, error) {
	addr := fmt.Sprintf(
		"%s:%d",
		config.Global.GetString("redis.host"),
		config.Global.GetInt("redis.port"),
	)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("error unable to establish Redis connection: %v", err)
	}

	return &Redix{client}, nil
}

func (rdx Redix) KeyExists(ctx context.Context, key string) (bool, error) {
	exists, err := rdx.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

func (rdx Redix) Get(ctx context.Context, key string, dst interface{}) error {
	res := rdx.client.HGetAll(ctx, key)
	if res.Err() != nil && !errors.Is(res.Err(), redis.Nil) {
		return res.Err()
	}

	if err := res.Scan(dst); err != nil {
		return err
	}

	return nil
}

func (rdx Redix) Set(ctx context.Context, key string, val interface{}) error {
	if err := rdx.client.HSet(ctx, key, val).Err(); err != nil {
		return err
	}

	return nil
}

func (rdx Redix) Close() error {
	return rdx.client.Close()
}
