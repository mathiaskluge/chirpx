package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerUpdateUser(w http.ResponseWriter, req *http.Request) {
	// Get token
	tokenString, err := utils.GetTokenFromRequest(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Get & Validate Token
	token, err := auth.ValidateJWT(tokenString, config.Env.JWTSecret)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: %w", err))
		return
	}

	// Get userID from Token
	userIDString, ok := token["userID"].(string)
	if !ok {
		utils.RespondWithError(w, http.StatusInternalServerError, errors.New("invalid or missing user id"))
		return
	}
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, errors.New("invalid or missing user id"))
		return
	}

	// Parse payload
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	// Check existance
	user, err := h.store.GetUserByID(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid ID"))
		return
	}

	// Hash new password (don't care for now if it's the same - gets replaced)
	// If this is relevant in the future, the users struct can be taken over from
	// GetUserByID and auth.ComparePasswords can be used to compare
	pwHash, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Update user

	if err := h.store.UpdateUser(userID, payload.Email, pwHash); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.CreateUserResponse{
		ID:         userID,
		Email:      payload.Email,
		IsUpgraded: user.IsUpgraded,
	})
}
