package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	opt, _ := redis.ParseURL(os.Getenv("DB_URL"))
	rdb := redis.NewClient(opt)

	return rdb
}
