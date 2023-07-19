package database

import (
	"TestTaskHaloLab/configuration"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func ConnectRedis(config *configuration.EnvConfigModel) *RedisCache {
	redisDB, err := strconv.Atoi(config.RedisDB)
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", config.RedisHost),
		Password: config.RedisPassword,
		DB:       redisDB,
	})
	err = client.Ping().Err()
	if err != nil {
		log.Fatal("Failed to connect to the redis!\n", err.Error())
	}
	log.Println("Successfully connected to the redis")
	return &RedisCache{client: client}
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(key, value, ttl).Err()
}
