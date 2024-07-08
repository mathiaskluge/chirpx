package chirp

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerDeleteChirp(w http.ResponseWriter, req *http.Request) {
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

	// Get chirpID from Request
	chirpID, err := strconv.Atoi(req.PathValue("chirpID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Get the chirp
	chirp, err := h.store.GetChirpByID(chirpID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Check if user is authorized to delete chirp
	if chirp.AuthorID != userID {
		utils.RespondWithError(w, http.StatusForbidden, errors.New("not authorized to delete chirp"))
		return
	}

	// Delete the chirp
	err = h.store.DeleteChirp(chirpID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
