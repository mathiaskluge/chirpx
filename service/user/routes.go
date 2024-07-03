package user

import (
	"fmt"
	"net/http"

	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users", h.handlerCreateUser)
}

func (h *Handler) handlerCreateUser(w http.ResponseWriter, req *http.Request) {
	// Parse Payload
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
	}
	// Check if User already exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	// Hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}
	// Generate a new user ID
	userID, err := h.store.GenerateUserID()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}
	// Create the Users
	err = h.store.CreateUser(types.User{
		ID:     userID,
		Email:  payload.Email,
		PwHash: hashedPassword,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	utils.RespondWithJSON(w, http.StatusCreated, nil)
}
