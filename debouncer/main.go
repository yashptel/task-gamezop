package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yashptel/debouncer/pkg/config"
	"github.com/yashptel/debouncer/pkg/debouncer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	conf := config.GetConfig()

	//  Initialize logger
	logger := initLogger()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Initialize redis client
	client, err := debouncer.NewClient(context.Background(), conf)
	if err != nil {
		zap.L().Fatal("error creating client", zap.Error(err))
	}
	defer client.Close()

	// Ping redis
	_, err = client.Ping()
	if err != nil {
		zap.L().Fatal("error pinging redis", zap.Error(err))
	}
	zap.L().Info("redis ping successful")

	// Subscribe to redis expiration events
	client.SubscribeEx(func(id string) error {
		req, err := http.NewRequest(http.MethodGet, conf.API+"/reward?id="+id, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("api call failed with status code %d", resp.StatusCode)
		}

		zap.L().Info("got response", zap.String("status", resp.Status))
		return nil
	})
}

func initLogger() *zap.Logger {

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
