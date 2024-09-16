package main

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"
)

var logger zap.Logger

func main() {
	r := chi.NewRouter()

	logger = zap.Logger{}

	r.Get("/api/address/geocode", nil)
	r.Use(LoggerMiddleware)

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("use middleware")
	})
}
