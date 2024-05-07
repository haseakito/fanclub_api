package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hackgame-org/fanclub_api/internal/redis"
	"golang.org/x/sync/singleflight"
)

type Cache[T any] struct {
	client     redis.Client
	expiration time.Duration
	sfg        *singleflight.Group
}

func NewCache[T any](rdb redis.Client, expiration time.Duration) *Cache[T] {
	return &Cache[T]{
		client:     rdb,
		expiration: expiration,
		sfg:        &singleflight.Group{},
	}
}

// GetOrSet tries to retrieve a cached item, and if it does not exist, it generates the item, caches, and returns it.
func (c *Cache[T]) GetOrSet(
	ctx context.Context,
	key string,
	callback func(context.Context) (T, error),
) (T, error) {
	// Use singleflight to ensure that only one load operation happens per key.
	res, err, _ := c.sfg.Do(key, func() (any, error) {
		// Attempt to get the cached value.
		bytes, exist, err := c.client.Get(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("error retrieving key %s from cache: %v", key, err)
		}
		if exist {
			return bytes, nil
		}

		// If cache did not exist, use the callback to generate the value.
		t, err := callback(ctx)
		if err != nil {
			return nil, err
		}

		// Marshal the result to store it in the cache.
		bytes, err = json.Marshal(t)
		if err != nil {
			return nil, fmt.Errorf("json marshal failed: %v", err)
		}

		// Set the new value in the cache with the specified expiration.
		if err := c.client.Set(ctx, key, bytes, c.expiration); err != nil {
			return nil, fmt.Errorf("error setting key %s in cache: %v", key, err)
		}

		return bytes, nil
	})

	var t T

	//
	if err != nil {
		return t, err
	}

	// Assert the type of the result
	bytes, ok := res.([]byte)
	if !ok {
		return t, fmt.Errorf("invalid type from cache result; expected []byte, got %T", res)
	}

	// Unmarshal the bytes into the expected type T.
	if err = json.Unmarshal(bytes, &t); err != nil {
		return t, fmt.Errorf("json unmarshal failed: %v", err)
	}

	return t, nil
}
