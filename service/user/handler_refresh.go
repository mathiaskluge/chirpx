package user

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	// Token is correct
	tokenString, err := utils.GetTokenFromRequest(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Session exists
	session, err := h.store.GetSession(tokenString)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Session is valid
	if err := auth.ValidateSession(session); err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Generate new acces token
	newToken, err := auth.CreateJWT(config.Env.JWTSecret, session.UserID, 60*60)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with access token
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"token": newToken,
	})
}
