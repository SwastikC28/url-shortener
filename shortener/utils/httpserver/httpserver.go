package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

var server *http.Server = nil

func StartServer() {
	logger := zerolog.Ctx(context.Background()).With().Str("module", "httpserver").Logger()

	server = &http.Server{}

	logger.Info().Msg("Starting server at PORT")
}

func StopServer(ctx context.Context) {
	logger := zerolog.Ctx(ctx).With().Str("module", "httpserver").Logger()

	if server != nil {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctxWithTimeout)
		if err != nil {
			return
		}

		logger.Info().Msg("Server shutdown successfully.")
	}
}
