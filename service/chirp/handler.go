package chirp

import (
	"net/http"
	"strconv"

	"github.com/mathiaskluge/chirpx/types"
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

func (h *Handler) handlerCreateChirp(w http.ResponseWriter, req *http.Request) {
	var payload types.CreateChirpPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
	}

	validatedBody, err := ValidateChirp(payload.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
	}

	chirpID, err := h.store.GenerateChirpID()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	err = h.store.CreateChirp(types.Chirp{
		ID:   chirpID,
		Body: validatedBody,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	utils.RespondWithJSON(w, http.StatusCreated, nil)
}
