package api

import (
	"log"
	"net/http"

	"github.com/mathiaskluge/chirpx/service/user"
)

type APIServer struct {
	addr string
	// db *db
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	mainRouter := http.NewServeMux()
	subrouter := http.NewServeMux()
	mainRouter.Handle("/api/", http.StripPrefix("/api", subrouter))

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, mainRouter)
}
