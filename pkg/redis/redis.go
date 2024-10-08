package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// RedisClient holds the Redis client
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient creates a new Redis client
func NewRedisClient(address string, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	return &RedisClient{Client: client}
}

// Ping sends a ping to the Redis server
func (r *RedisClient) Ping() error {
	pong, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("error pinging Redis server: %s", err)
	}
	if pong != "PONG" {
		return fmt.Errorf("invalid models received from Redis server: %s", pong)
	}

	return nil
}

// Set sets a key-value pair in Redis
func (r *RedisClient) Set(key string, value interface{}, ttl ...int) error {
	var expiration time.Duration
	if len(ttl) > 0 {
		expiration = time.Duration(ttl[0]) * time.Second
	} else {
		expiration = 0
	}
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("error setting value in Redis: %s", err)
	}

	return nil
}

// Get retrieves the value of a key from Redis
func (r *RedisClient) Get(key string) (string, error) {
	value, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist in Redis", key)
	} else if err != nil {
		return "", fmt.Errorf("error getting value from Redis: %s", err)
	}

	return value, nil
}

// Del deletes one or more keys from Redis
func (r *RedisClient) Del(keys ...string) error {
	err := r.Client.Del(ctx, keys...).Err()
	if err != nil {
		return fmt.Errorf("error deleting keys from Redis: %s", err)
	}

	return nil
}
