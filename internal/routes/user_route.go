package routes

import (
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/handlers"
	"github.com/sammydepoj/golang-rest-api/internal/middlewares"
)

func SetUpUserRoute(mux *http.ServeMux, handler *handlers.Handler) {
	userMux := http.NewServeMux()

	userMux.HandleFunc("POST /register", handler.CreateUserHandler())
	userMux.HandleFunc("POST /login", handler.LoginUserHandler())
	userMux.Handle("GET /profile", middlewares.AuthMiddleWare(http.HandlerFunc(handler.UserProfileHandler())))
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

}
