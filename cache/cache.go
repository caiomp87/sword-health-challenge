package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/caiomp87/sword-health-challenge/interfaces"
	"github.com/go-redis/redis/v8"
)

var CacheService interfaces.ICache

type cacheHelper struct {
	client *redis.Client
}

func NewCache() interfaces.ICache {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DB:           0,
		MaxRetries:   3,
		WriteTimeout: time.Second * 5,
	})

	return &cacheHelper{
		client,
	}
}

func (c *cacheHelper) Publish(ctx context.Context, key string, payload []byte) error {
	publisher := c.client.Publish(ctx, key, payload)
	if err := publisher.Err(); err != nil {
		return err
	}

	return nil
}

func (c *cacheHelper) Subscribe(ctx context.Context, key string) error {
	subscriber := c.client.Subscribe(ctx, key)

	var notification string
	for {
		message, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(message.Payload), &notification); err != nil {
			return err
		}

		log.Println("Notification received:", notification)
	}
}

func (c *cacheHelper) Ping(ctx context.Context) error {
	ping := c.client.Ping(ctx)

	if ping.Err() != nil || ping.Val() != "PONG" {
		return fmt.Errorf("could not ping the redis server: %v", ping.Err())
	}

	return nil
}
