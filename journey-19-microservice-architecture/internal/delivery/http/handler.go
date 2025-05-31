package http

import (
	"encoding/json"
	"net/http"

	"microservice-architecture/internal/logger"
	"microservice-architecture/internal/usecase"
)

type UserHandler struct {
	UC *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UC: uc}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")

	user, err := h.UC.GetUserByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	logger.Log.WithField("user_id", id).Info("User fetched successfully")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
