package redis

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"resedist/pkg/config"
)

func Connect() {
	cfg := config.Get()

	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password, // No password set
		DB:       cfg.Redis.DB,       // Use default DB
		//Protocol: 2,  // Connection protocol
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(Client); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(Client); err != nil {
		panic(err)
	}

	ctx := context.Background()
	err := Client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
}
