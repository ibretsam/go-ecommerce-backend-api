package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		checkErrorPanic(err, "Redis initialization failed")
	}

	fmt.Println("Redis init success")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Set failed", zap.Error(err))
		panic(err)
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Get failed", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("score", zap.String("score", value))
}
