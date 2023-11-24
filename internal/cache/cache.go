package cache

import (
	"context"
	"encoding/json"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RDBLayer struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache
type Caching interface {
	Set(ctx context.Context, jobID uint, jobData models.Jobs) error
	Get(ctx context.Context, jobID uint) (string, error)
	VerficationCodeSet(ctx context.Context,  email string,verficationCode int) error
	VerficationCodeGet(ctx context.Context, email string)(string, error)
}

func NewRDBLayer(rdb *redis.Client) Caching {
	return &RDBLayer{
		rdb: rdb,
	}
}

func (c *RDBLayer) Set(ctx context.Context, jobID uint, jobData models.Jobs) error {
	jid := strconv.FormatUint(uint64(jobID), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}

	err = c.rdb.Set(ctx, jid, val, 5*time.Minute).Err()
	return err

}

func (c *RDBLayer) Get(ctx context.Context, jobID uint) (string, error) {
	jid := strconv.FormatUint(uint64(jobID), 10)
	str, err := c.rdb.Get(ctx, jid).Result()
	return str, err
}
func (c* RDBLayer) VerficationCodeSet(ctx context.Context, email string,verficationCode int) error{
	err := c.rdb.Set(ctx, email, verficationCode, 5*time.Minute).Err()
	return err
}

func(c* RDBLayer) VerficationCodeGet(ctx context.Context, email string)(string, error){
	str, err := c.rdb.Get(ctx, email).Result()
	return str, err
}
