package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

var server *http.Server = nil

func StartServer(router *mux.Router) {
	logger := zerolog.Ctx(context.Background()).With().Str("module", "httpserver").Logger()
	config := newhttpserverconfig()

	server = &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Port),
		Handler: CORSMiddleware(router),
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

func CORSMiddleware(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	return c.Handler(next)
}
