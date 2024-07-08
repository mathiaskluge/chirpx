package chirp

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerCreateChirp(w http.ResponseWriter, req *http.Request) {
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

	// Parse and validate request body
	var payload types.CreateChirpPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
	}
	validatedBody, err := ValidateChirp(payload.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
	}

	// Generate a chirp ID
	chirpID, err := h.store.GenerateChirpID()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	// Create the chirp in the database
	err = h.store.CreateChirp(types.Chirp{
		ID:       chirpID,
		Body:     validatedBody,
		AuthorID: userID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	utils.RespondWithJSON(w, http.StatusCreated, types.CreateChirpResponse{
		ID:       chirpID,
		Body:     validatedBody,
		AuthorID: userID,
	})
}
