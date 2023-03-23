package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

var (
	rclient *redis.Client
)

func init() {
	rclient = redis.NewClient(&redis.Options{
		Addr:     "106.55.104.223:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func TestRD(t *testing.T) {
	ctx := context.Background()

	err := rclient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rclient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rclient.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

