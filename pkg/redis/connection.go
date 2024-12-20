package redis

import (
	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	return Client
}
