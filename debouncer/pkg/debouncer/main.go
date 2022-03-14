package debouncer

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/yashptel/debouncer/pkg/config"
	"go.uber.org/zap"
)

type Debouncer struct {
	ctx         context.Context
	conf        config.AppConfig
	redisClient *redis.Client
}

type DebouncerClient interface {
	Ping() (string, error)
	Close() error
	SubscribeEx(func(msg string) error)
}

func NewClient(ctx context.Context, conf config.AppConfig) (DebouncerClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Host,
		Password: conf.Redis.Pass,
		DB:       conf.Redis.Db,
	})

	return &Debouncer{
		ctx:         ctx,
		conf:        conf,
		redisClient: client,
	}, nil
}

func (d *Debouncer) Ping() (string, error) {
	return d.redisClient.Ping(d.ctx).Result()
}

func (d *Debouncer) Close() error {
	return d.redisClient.Close()
}

func (d *Debouncer) SubscribeEx(callback func(msg string) error) {

	d.redisClient.ConfigSet(d.ctx, "notify-keyspace-events", "Ex")

	pubsub := d.redisClient.Subscribe(d.ctx, "__keyevent@0__:expired")
	defer pubsub.Close()

	zap.L().Info("Subscribed to redis expiration events. Waiting for events...")

	for msg := range pubsub.Channel() {
		zap.L().Info("received message", zap.String("channel", msg.Channel), zap.String("payload", msg.Payload))
		err := callback(msg.Payload)
		if err != nil {
			zap.L().Error("error in callback", zap.Error(err))
		}
	}
}
