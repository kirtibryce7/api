package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx    = context.Background()
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

func RedisCon() *redisClient {
	LoadEnv()
	c := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_ip"),
		Password: "",
		DB:       0,
	})
	pong, err := c.Ping(ctx).Result()
	fmt.Println("pong", pong)
	if err != nil {
		fmt.Println(err)
	}
	client.c = c
	return client
}

func (client *redisClient) GetKey(key string, src interface{}) error {
	val, err := client.c.Get(ctx, key).Result()
	if err == redis.Nil || err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (client *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	err = client.c.Set(ctx, key, cacheEntry, expiration).Err()
	if err != nil {
		fmt.Println(err)
	}
	return nil

}
func (client *redisClient) DelKey(key string) error {

	var foundedRecordCount int = 0
	iter := client.c.Scan(ctx, 0, key, 0).Iterator()
	fmt.Printf("Trying to Remove from Redish for Refresh Token= %s\n", key)
	for iter.Next(ctx) {
		fmt.Printf("Deleted= %s\n", iter.Val())
		client.c.Del(ctx, iter.Val())
		foundedRecordCount++
	}
	if err := iter.Err(); err != nil {
		return err
	}
	fmt.Printf("Deleted Count %d\n", foundedRecordCount)
	return nil
}
