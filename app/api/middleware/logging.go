package middleware

import (
	"comies/app/telemetry"
	"net/http"

	"go.uber.org/zap"
)

type LoggerContextKey struct{}

func Logging(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			handler.ServeHTTP(writer, request.WithContext(
				telemetry.SetLoggerToContext(
					request.Context(),
					logger.With(
						zap.String("method", request.Method),
						zap.String("path", request.URL.Path),
					),
				),
			))
		})
	}
}
