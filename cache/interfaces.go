package cache

import (
	"context"
	"fmt"
	"time"
)

var (
	ErrCacheMiss = func(key string) error {
		return fmt.Errorf("cache miss for key: %s.\nrecommended action: load from persistent data source", key)
	}
	ErrDefaultCache = fmt.Errorf("the requested action could not be completed due to an internal error. please try again later")
)

type ICache interface {
	// GetString - get a string/JSON value
	GetString(context.Context, string) (string, error)

	// GetHash - get a hash of values (e.g. a user object)
	GetHash(context.Context, string) (map[string]interface{}, error)

	// GetList - get a list of values (e.g. a list of countries etc.)
	GetList(context.Context, string) ([]any, error)

	// StoreString - store a string/JSON value
	StoreString(context.Context, string, string, time.Duration) error

	// StoreHash - store a hash of values (e.g. a user object)
	StoreHash(context.Context, string, map[string]interface{}, time.Duration) error

	// StoreList - store a list of values (e.g. a list of countries etc.)
	StoreList(context.Context, string, []any, time.Duration) error

	Delete(context.Context, string) error
}
