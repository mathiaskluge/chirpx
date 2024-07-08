package user

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	token, err := utils.GetTokenFromRequest(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	session, err := h.store.GetSession(token)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	session.Revoked = true

	err = h.store.UpdateSession(token, session)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusNoContent, nil)

}
