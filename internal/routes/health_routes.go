package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/handlers"
)

func setupHealthRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
