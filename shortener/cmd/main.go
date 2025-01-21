package main

import (
	"context"

	"github.com/rs/zerolog"
)

var serviceName = "url-shortener"

func main() {
	ctx := context.Background()

	logger := zerolog.Ctx(ctx).With().Str("service", serviceName).Logger()
	logger.Info().Msg("Starting microservice")
}
