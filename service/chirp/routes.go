package chirp

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/types"
)

type Handler struct {
	store types.ChirpStore
}

func NewHandler(store types.ChirpStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /chirps", h.handlerGetChirps)
	router.HandleFunc("GET /chirps/{chirpID}", h.handlerGetChirpByID)
	router.HandleFunc("POST /chirps", h.handlerCreateChirp)
}
