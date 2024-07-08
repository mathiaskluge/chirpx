package chirp

import (
	"net/http"
	"strconv"

	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := h.store.GetChirps()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, chirps)
}

func (h *Handler) handlerGetChirpByID(w http.ResponseWriter, req *http.Request) {
	chirpID, err := strconv.Atoi(req.PathValue("chirpID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	chirp, err := h.store.GetChirpByID(chirpID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, chirp)
}
