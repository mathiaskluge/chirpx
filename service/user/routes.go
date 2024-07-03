package user

import (
	"net/http"

	"github.com/mathiaskluge/chirpx/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users", h.handlerCreateUser)
}
