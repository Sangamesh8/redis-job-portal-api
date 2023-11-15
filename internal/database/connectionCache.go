package database

import "github.com/redis/go-redis/v9"

func RedisConnection() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost : 8080",
		Password: "",
		DB: 0,
		Protocol: 3,
	})
	return rdb
}