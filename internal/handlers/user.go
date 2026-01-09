package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sammydepoj/golang-rest-api/internal/dtos"
	"github.com/sammydepoj/golang-rest-api/internal/store"
	"github.com/sammydepoj/golang-rest-api/internal/util"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create context
		ctx := r.Context()

		//user request body

		var request dtos.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		hashedPassword, err := util.HashPassword(request.Password)

		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		_, err = h.Queries.CreateUsers(ctx, store.CreateUsersParams{
			Username: request.Username,
			Email:    request.Email,
			Password: hashedPassword,
		})
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}

		util.RespondWithSuccess(w, http.StatusCreated, "User created successfully", request.Username)

	}
}
