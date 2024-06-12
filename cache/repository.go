package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// cache is a private struct that implements ICache
type cache struct {
	ICache
}

func New() ICache {
	if GetCacheClient() == nil {
		log.Fatalln("cache client is not initialized")
	}
	return &cache{}
}

func (c *cache) GetString(ctx context.Context, key string) (string, error) {
	data, err := GetCacheClient().Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", ErrCacheMiss(key)
		}
		return "", ErrDefaultCache
	}

	return data, nil
}

func (c *cache) StoreString(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := GetCacheClient().Set(ctx, key, value, expiration).Err()
	if err != nil {
		return ErrDefaultCache
	}
	return nil
}
