package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sammydepoj/golang-rest-api/internal/auth"
	"github.com/sammydepoj/golang-rest-api/internal/dtos/requests"
	"github.com/sammydepoj/golang-rest-api/internal/errorhandler"
	"github.com/sammydepoj/golang-rest-api/internal/middlewares"
	"github.com/sammydepoj/golang-rest-api/internal/store"
	"github.com/sammydepoj/golang-rest-api/internal/successresponse"
	"github.com/sammydepoj/golang-rest-api/internal/util"
	"github.com/sammydepoj/golang-rest-api/internal/validation"
)

//login

func (h *Handler) LoginUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create context
		ctx := r.Context()

		//user request body
		var request requests.LoginUserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			errorhandler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		//validate the request
		if err := validation.Validate(&request); err != nil {
			errorhandler.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		// fetch the user from DB using the store queries
		user, err := h.Queries.GetUserByUsernameOrEmail(ctx, store.GetUserByUsernameOrEmailParams{
			Username: request.Username,
			Email:    request.Username,
		})
		if err != nil {
			errorhandler.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		// compare the password
		if !util.ComparePassword(user.Password, request.Password) {
			errorhandler.RespondWithError(w, http.StatusInternalServerError, "Invalid email or password")
			return
		}
		// generate jwt token
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		token, err := auth.GenerateJwt(int64(user.ID), user.Username, jwtKey)
		if err != nil {
			errorhandler.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}
		successresponse.RespondWithSuccess(w, http.StatusOK, "Login successful", map[string]string{"token": token})
	}
}

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create context
		ctx := r.Context()

		//user request body

		var request requests.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			errorhandler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		log.Println(request)
		//validate the request
		if err := validation.Validate(&request); err != nil {
			errorhandler.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		hashedPassword, err := util.HashPassword(request.Password)

		if err != nil {
			errorhandler.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		_, err = h.Queries.CreateUsers(ctx, store.CreateUsersParams{
			Username: request.Username,
			Email:    request.Email,
			Password: hashedPassword,
			Created:  sql.NullTime{Time: time.Now(), Valid: true},
			Updated:  sql.NullTime{Time: time.Now(), Valid: true},
		})
		if err != nil {
			log.Printf("Failed to create user: %v", err)
			errorhandler.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}

		successresponse.RespondWithSuccess(w, http.StatusCreated, "User created successfully", request.Username)

	}
}

func (h *Handler) UserProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		claims, ok := ctx.Value(middlewares.UserClaimsKey).(*auth.Claims)
		if !ok {
			errorhandler.RespondWithError(w, http.StatusBadRequest, " Please login to continue")
			return
		}
		userID := claims.UserID

		user, err := h.Queries.GetUserByID(ctx, int32(userID))
		if err != nil {
			errorhandler.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		successresponse.RespondWithSuccess(w, http.StatusOK, "User profile fetched successfully", user)

	}

}
