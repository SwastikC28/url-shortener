package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var server *http.Server = nil

func StartServer(router *mux.Router) {
	logger := zerolog.Ctx(context.Background()).With().Str("module", "httpserver").Logger()
	config := newhttpserverconfig()

	server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Port),
		Handler: router,
	}

	logger.Info().Msgf("Starting server at port %d", config.Port)
	if err := server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
		logger.Err(err).Msg("Error starting server")
		os.Exit(1)
	}
}

func StopServer(ctx context.Context) {

	if server != nil {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err := server.Shutdown(ctxWithTimeout)
		if err != nil {
			return
		}

	}
}
