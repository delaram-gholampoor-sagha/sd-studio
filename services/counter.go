package services

import (
	"context"

	"github.com/delaram-gholampoor-sagha/sd-studio/storage"
)

type CounterService struct {
	redis *storage.RedisStorage
}

func NewCounterService(redis *storage.RedisStorage) *CounterService {
	return &CounterService{redis: redis}
}

func (c *CounterService) Increment(ctx context.Context, key string) error {
	return c.redis.IncrementCounter(ctx, key)
}

func (c *CounterService) GetCounter(ctx context.Context, key string) (int64, error) {
	return c.redis.GetCounter(ctx, key)
}
