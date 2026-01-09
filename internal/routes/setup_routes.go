package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	setupHealthRoutes(mux, handler)
}
