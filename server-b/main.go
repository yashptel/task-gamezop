package main

import (
	"context"

	"github.com/yashptel/server-b/pkg/http"
	"github.com/yashptel/server-b/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// Run a new HTTP server
	http.RunHttpServer(context.Background())
}
