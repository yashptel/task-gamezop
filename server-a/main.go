package main

import (
	"context"

	"github.com/yashptel/go-api-template/pkg/http"
	"github.com/yashptel/go-api-template/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Run a new HTTP server
	http.RunHttpServer(context.Background())
}
