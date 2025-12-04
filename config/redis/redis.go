package configs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisConfig() *redis.Client {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-14126.c295.ap-southeast-1-1.ec2.cloud.redislabs.com:14126",
		Username: "default",
		Password: "hkWRmHvIgB9QhsTRUqvBuSK2uODszhwT",
		DB:       0,
	})
	fmt.Println(rdb.Ping(ctx))
	return rdb
}
