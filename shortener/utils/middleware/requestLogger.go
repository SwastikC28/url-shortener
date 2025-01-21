package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

func ReqLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if r.Header.Get("X-Correlation-Id") == "" {
				r.Header.Set("X-Correlation-Id", uuid.NewV4().String())
			}

			logger := zerolog.Ctx(r.Context()).With().Str("Correlation-Id", r.Header.Get("X-Correlation-Id")).Logger()

			startTime := time.Now()

			// Log the details
			logger.Info().
				Str("METHOD", r.Method).
				Str("URL", r.URL.String()).
				Msg("Begin")

			next.ServeHTTP(w, r)

			timeElapsed := time.Since(startTime).Milliseconds()
			timeElapsedStr := strconv.FormatInt(timeElapsed, 10)

			logger.Info().
				Str("METHOD", r.Method).
				Str("URL", r.URL.String()).
				Str("Time-Elapsed", timeElapsedStr).
				Msg("End")
		},
	)
}
