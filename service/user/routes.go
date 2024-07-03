package user

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /test", h.handleStuff)
}

func (h *Handler) handleStuff(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("this is the get test stuff"))
}
