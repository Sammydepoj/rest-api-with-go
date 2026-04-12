package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/errorhandler"
	"github.com/sammydepoj/golang-rest-api/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetUpUserRoute(mux, handler)
	setupHealthRoutes(mux, handler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		errorhandler.RespondWithNotFound(w)
	})
}
