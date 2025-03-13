package store

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore() Store {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	return &redisStore{
		client,
	}
}

func (rs *redisStore) Set(key string, value int, expiry time.Duration) error {
	ctx := context.Background()
	res := rs.client.Set(ctx, key, value, expiry)
	if _, err := res.Result(); err != nil {
		return err
	}
	return nil
}

func (rs *redisStore) Get(key string) (int, error) {
	ctx := context.Background()
	res := rs.client.Get(ctx, key)
	resStr, err := res.Result()
	if err != nil {
		return 0, err
	}
	op, err := strconv.ParseInt(resStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(op), nil
}

func (rs *redisStore) GetTTL(key string) (time.Duration, error) {
	ctx := context.Background()
	res := rs.client.TTL(ctx, key)
	duration, err := res.Result()
	if err != nil {
		return 0, err
	}
	return duration, nil
}

func (rs *redisStore) Increment(key string) (int, error) {
	ctx := context.Background()
	res := rs.client.Incr(ctx, key)
	op, err := res.Result()
	if err != nil {
		return 0, err
	}
	return int(op), nil
}

func (rs *redisStore) Decrement(key string) (int, error) {
	ctx := context.Background()
	res := rs.client.Decr(ctx, key)
	op, err := res.Result()
	if err != nil {
		return 0, err
	}
	return int(op), nil
}

func (rs *redisStore) Delete(key string) error {
	ctx := context.Background()
	res := rs.client.Del(ctx, key)
	_, err := res.Result()
	if err != nil {
		return err
	}
	// TODO: Parse opcode and pass errors accordingly
	return nil
}

func (rs *redisStore) SetEx(key string, expiry time.Duration) error {
	ctx := context.Background()
	res := rs.client.SetEx(ctx, key, nil, expiry)
	_, err := res.Result()
	if err != nil {
		return err
	}
	return nil
}
