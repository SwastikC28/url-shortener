package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/internal/config"
	"url-shortener/utils/httpserver"
	"url-shortener/utils/routing"

	_ "url-shortener/utils/logging"

	"github.com/rs/zerolog"
)

var serviceName = "url-shortener"

func main() {
	ctx := context.Background()

	logger := zerolog.Ctx(ctx).With().Str("service", serviceName).Logger()
	logger.Info().Msg("Starting microservice")

	router := routing.NewDefaultRouter()
	config.InitializeApp(router)

	go httpserver.StartServer(router)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	httpserver.StopServer(ctx)
	logger.Info().Msg("Server shutdown successfully")
	os.Exit(0)
}
