package user

import (
	"fmt"
	"net/http"

	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerUpgradeUser(w http.ResponseWriter, req *http.Request) {
	// Parse payload
	var payload types.UpgradeUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	if payload.Event != "user.upgraded" {
		err := fmt.Errorf("invalid event type: %s", payload.Event)
		utils.RespondWithError(w, http.StatusNoContent, err)
		return
	}

	// check user id
	_, err := h.store.GetUserByID(payload.Data.UserID)
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
