package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sammydepoj/golang-rest-api/internal/auth"
	"github.com/sammydepoj/golang-rest-api/internal/errorhandler"
)

type contextKey string

const UserClaimsKey contextKey = "user_claims"

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errorhandler.RespondWithError(w, http.StatusUnauthorized, "No token provided")
			return
		}
		// Bearer token
		// tokenString := authHeader[len("Bearer "):]
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				errorhandler.RespondWithError(w, http.StatusUnauthorized, "Invalid token signature")
				return
			}
			errorhandler.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		if token.Valid {
			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		} else {
			errorhandler.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

	})
}
