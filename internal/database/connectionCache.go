package database

import (
	"context"
	"job-portal-api/config"

	"github.com/redis/go-redis/v9"
)

func RedisConnection(cfg config.Config) (*redis.Client,error ){
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisConfig.RedisAddr,
		Password: cfg.RedisConfig.RedisPassword,
		DB:       cfg.RedisConfig.RedisDb,
	})
	_,err :=rdb.Ping(context.Background()).Result()
	if err != nil{
		return nil, err
	}
	return rdb,nil
}
