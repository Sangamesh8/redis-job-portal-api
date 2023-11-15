package cache

import "github.com/redis/go-redis/v9"

type RDBLayer struct {
	rdb *redis.NewClient
}

type Caching interface{
	
}