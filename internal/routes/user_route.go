package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/handlers"
)

func SetUpUserRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("POST /user/register", handler.CreateUserHandler())

}
