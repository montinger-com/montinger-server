package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/montinger-com/montinger-server/config"
	"github.com/montinger-com/montinger-server/pkg/redis"
	"github.com/rashintha/logger"
)

var redisClient *redis.RedisClient

var tryingReconnect = false
var connected = false

func Init() {
	logger.Defaultln("Initializing Cache")

	redisClient = redis.NewRedisClient(fmt.Sprintf("%v:%v", config.REDIS_HOST, config.REDIS_PORT), config.REDIS_PASS, config.REDIS_DB)
	go periodicallyCheckRedis()
}

func periodicallyCheckRedis() {
	checkInterval := 30 * time.Second // Adjust the interval as needed
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	if !tryingReconnect && !connected {
		err := redisClient.Ping()
		if err == nil {
			logger.Defaultln("Redis cache connected!")
		} else {
			logger.Errorln(err.Error())
		}
	}

	for {
		select {
		case <-ticker.C:
			err := redisClient.Ping()
			if err != nil {
				if connected {
					logger.Errorln("Redis cache connection lost. Reconnecting...")
					connected = false
					tryingReconnect = true
					// Reconnect logic goes here
					redisClient = redis.NewRedisClient(fmt.Sprintf("%v:%v", config.REDIS_HOST, config.REDIS_PORT), config.REDIS_PASS, config.REDIS_DB)
				}
			} else {
				if !connected && tryingReconnect {
					logger.Defaultln("Redis cache reconnected!")
					connected = true
					tryingReconnect = false
					continue
				}

			}
		}
	}
}

func Set(key string, value interface{}, ttl ...int) error {
	var data []byte
	var err error
	switch v := value.(type) {
	case string:
		data = []byte(v)
	default:
		data, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}
	return redisClient.Set(key, data, ttl...)
}

func Get(key string) (interface{}, error) {
	data, err := redisClient.Get(key)
	if err != nil {
		return "", err
	}
	var cachedData interface{}
	if err = json.Unmarshal([]byte(data), &cachedData); err != nil {
		cachedData = string(data)
	}
	return cachedData, nil
}

func Delete(keys ...string) error {
	return redisClient.Del(keys...)
}
