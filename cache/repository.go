package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// cache is a private struct that implements ICache
type cache struct {
	client *redis.Client
	ICache
}

func New(client *redis.Client) ICache {
	return &cache{client: client}
}

func (c *cache) GetString(ctx context.Context, key string) (string, error) {
	data, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrCacheMiss(key)
		}
		return "", ErrDefaultCache
	}

	return data, nil
}

func (c *cache) StoreString(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return ErrDefaultCache
	}
	return nil
}
