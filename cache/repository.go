package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// cache is a private struct that implements ICache
type cache struct {
	client *redis.Client
	ICache
}

func NewCache(client *redis.Client) ICache {
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

func (c *cache) StoreList(ctx context.Context, key string, values []*interface{}, expiration time.Duration) error {
	if err := c.client.SAdd(ctx, key, toJSONArray(values)).Err(); err != nil {
		return ErrDefaultCache
	}

	return nil
}

func (c *cache) GetList(ctx context.Context, key string) ([]*interface{}, error) {
	data, err := c.client.SMembers(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrCacheMiss(key)
		}
		return nil, ErrDefaultCache
	}

	if len(data) == 0 {
		return nil, ErrCacheMiss(key)
	}

	result := fromJSONArray(data)
	return result, nil
}

func toJsonString(data *interface{}) string {
	encoded, err := json.Marshal(&data)
	if err != nil {
		return ""
	}

	return string(encoded)
}

func toJSONArray(data []*interface{}) []string {
	var result []string
	for _, v := range data {
		result = append(result, toJsonString(v))
	}

	return result
}

func fromJSONArray(data []string) []*interface{} {
	var result []*interface{}
	for _, v := range data {
		var temp interface{}
		if err := json.Unmarshal([]byte(v), &temp); err != nil {
			continue
		}
		result = append(result, &temp)
	}

	return result
}
