package chirp

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

type Handler struct {
	store types.ChirpStore
}

func NewHandler(store types.ChirpStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /chirps", h.handlerGetChirps)
	router.HandleFunc("GET /chirps/{chirpsID}", h.handlerGetChirpByID)
	router.HandleFunc("POST /chirps", h.handlerCreateChirp)
}

func (h *Handler) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	chirps, err := h.store.GetChirps()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}
	utils.RespondWithJSON(w, http.StatusOK, chirps)
}
func (h *Handler) handlerGetChirpByID(w http.ResponseWriter, req *http.Request) {

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
