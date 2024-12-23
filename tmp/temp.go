package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"resedist/pkg/converters"
	"time"
)

var rdb *redis.Client

func getOrSetWithTag(ctx context.Context, key, tag string, ttl time.Duration, fetchFunc func() (string, error)) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key not found, fetch new data
		val, err = fetchFunc()
		if err != nil {
			return "", err
		}
		// Store data in cache
		err = rdb.Set(ctx, key, val, ttl).Err()
		if err != nil {
			return "", err
		}
		// Associate the key with the tag
		err = rdb.SAdd(ctx, "tag:"+tag, key).Err()
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return val, nil
}
func invalidateTag(ctx context.Context, tag string) error {
	keys, err := rdb.SMembers(ctx, "tag:"+tag).Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		rdb.Del(ctx, key)
	}
	return rdb.Del(ctx, "tag:"+tag).Err()
}

// cacheWithTag caches a value with an associated tag
func cacheWithTag(ctx context.Context, key, tag string, value string, ttl time.Duration) error {
	err := rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}
	return rdb.SAdd(ctx, "tag:"+tag, key).Err()
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,
		//Protocol: 2,  // Connection protocol
	})
	ctx := context.Background()

	key := "product:456"
	tag := "products"
	ttl := 10 * time.Minute

	fetchFunc := func() (string, error) {
		fmt.Println("Fetching product data...")
		return "Product Data for ID 456", nil
	}

	// Get or set the value with tag support
	value, err := getOrSetWithTag(ctx, key, tag, ttl, fetchFunc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Cached Value:", value)

	// Invalidate all cache associated with the tag
	fmt.Println("Invalidating tag:", tag)
	err = invalidateTag(ctx, tag)
	if err != nil {
		fmt.Println("Error invalidating tag:", err)
		return
	}
	fmt.Println("Tag invalidated successfully!")
}

func main22() {
	var UrlData = make(map[string][]string)

	key := "title"
	key2 := "title2"
	value := "body"
	value2 := "body2"

	UrlData = map[string][]string{
		key:  {value, value2},
		key2: {value, value2},
	}

	fmt.Println(UrlData)

	str := converters.UrlValuesToString(UrlData)
	fmt.Println(str)

	UrlData2 := converters.StringToUrlValues(str)
	fmt.Println(UrlData2)

	str = converters.UrlValuesToString(UrlData2)
	fmt.Println(str)

}

func hasKeyValue(m map[string][]string, key string, value string) bool {
	if v, exists := m[key]; exists {
		for _, val := range v {
			if val == value {
				return true
			}
		}
	}
	return false
}
