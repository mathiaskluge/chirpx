package user

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerUpgradeUser(w http.ResponseWriter, req *http.Request) {
	// Get API Key from Header
	APIKey, err := utils.GetAPIKeyFromRequest(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	if APIKey != config.Env.PolkaSecret {
		utils.RespondWithError(w, http.StatusUnauthorized, err)
		return
	}

	// Parse payload
	var payload types.UpgradeUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Event != "user.upgraded" {
		utils.RespondWithJSON(w, http.StatusNoContent, nil)
		return
	}

	// check user id
	_, err = h.store.GetUserByID(payload.Data.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err)
		return
	}

	err = h.store.UpgradeUser(payload.Data.UserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)
}
