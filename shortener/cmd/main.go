package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/utils/httpserver"
	_ "url-shortener/utils/logging"

	"github.com/rs/zerolog"
)

var serviceName = "url-shortener"

func main() {
	ctx := context.Background()

	logger := zerolog.Ctx(ctx).With().Str("service", serviceName).Logger()
	logger.Info().Msg("Starting microservice")

	go httpserver.StartServer(nil)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	httpserver.StopServer(ctx)
	logger.Info().Msg("Server shutdown successfully")
	os.Exit(0)
}
