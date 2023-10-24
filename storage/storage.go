package storage

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string, password string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return &RedisStorage{client: client}
}

func (r *RedisStorage) IncrementCounter(ctx context.Context, key string) error {
	_, err := r.client.Incr(ctx, key).Result()
	return err
}

func (r *RedisStorage) GetCounter(ctx context.Context, key string) (int64, error) {
	return r.client.Get(ctx, key).Int64()
}

func (r *RedisStorage) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisStorage) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
