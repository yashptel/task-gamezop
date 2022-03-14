package debounce

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yashptel/go-api-template/pkg/config"
	"go.uber.org/zap"
)

type DebounceClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type Debounce struct {
	ctx         context.Context
	redisClient *redis.Client
	conf        config.AppConfig
}

var client *redis.Client
var once sync.Once

func NewDebounce(ctx context.Context) DebounceClient {
	return &Debounce{
		ctx:         ctx,
		redisClient: getRedisClient(ctx),
		conf:        config.GetConfig(),
	}
}

func getRedisClient(ctx context.Context) *redis.Client {
	once.Do(func() {
		conf := config.GetConfig()

		client = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Host,
			Password: conf.Redis.Pass,
			DB:       conf.Redis.Db,
		})

		pong, err := client.Ping(ctx).Result()
		if err != nil {
			zap.L().Fatal("redis ping error", zap.Error(err))
		}
		zap.L().Info("redis ping", zap.String("pong", pong))

		client.ConfigSet(ctx, "notify-keyspace-events", "Ex")
	})
	return client
}

func (d *Debounce) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if config.IsTestEnv() {
		return nil
	}
	return d.redisClient.Set(ctx, key, value, expiration).Err()
}
