package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"resedist/pkg/applog"
	"time"
)

var Client *redis.Client

func GetOrSetWithTagJSON[T any](ctx context.Context, key, tag string, ttl time.Duration, fetchFunc func() (T, error)) (T, error) {
	var result T

	// Attempt to get the data from Redis
	cachedData, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key not found, fetch new data
		fetchedData, err := fetchFunc()
		if err != nil {
			return result, err
		}

		// Serialize data to JSON
		jsonData, err := json.Marshal(fetchedData)
		if err != nil {
			return result, err
		}

		// Store serialized data in Redis
		err = Client.Set(ctx, key, jsonData, ttl).Err()
		if err != nil {
			return result, err
		}

		// Associate the key with the tag
		err = Client.SAdd(ctx, "tag:"+tag, key).Err()
		if err != nil {
			return result, err
		}

		return fetchedData, nil
	} else if err != nil {
		return result, err
	}

	// Deserialize cached JSON data
	err = json.Unmarshal([]byte(cachedData), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetOrSetJSON[T any](ctx context.Context, key string, ttl time.Duration, fetchFunc func() (T, error)) (T, error) {
	var result T

	// Attempt to get the data from Redis
	cachedData, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key not found, fetch new data
		fetchedData, err := fetchFunc()
		if err != nil {
			return result, err
		}

		// Serialize data to JSON
		jsonData, err := json.Marshal(fetchedData)
		if err != nil {
			return result, err
		}

		// Store serialized data in Redis
		err = Client.Set(ctx, key, jsonData, ttl).Err()
		if err != nil {
			return result, err
		}

		return fetchedData, nil
	} else if err != nil {
		return result, err
	}

	// Deserialize cached JSON data
	err = json.Unmarshal([]byte(cachedData), &result)
	if err != nil {
		return result, err
	}

	applog.Info(fmt.Sprintf("Fetch %s from redis", key))
	return result, nil
}
