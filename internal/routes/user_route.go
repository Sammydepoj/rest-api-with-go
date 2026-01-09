package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/handlers"
)

func SetUpUserRoute(mux *http.ServeMux, handler *handlers.Handler) {
	userMux := http.NewServeMux()

	userMux.HandleFunc("POST /login", handler.LoginUserHandler())
	userMux.HandleFunc("POST /register", handler.CreateUserHandler())
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

}
