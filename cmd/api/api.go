package api

import (
	"log"
	"net/http"

	"github.com/mathiaskluge/chirpx/db"
	"github.com/mathiaskluge/chirpx/service/chirp"
	"github.com/mathiaskluge/chirpx/service/user"
)

type APIServer struct {
	addr string
	db   *db.DB
}

func NewAPIServer(addr string, db *db.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mainRouter := http.NewServeMux()
	subrouter := http.NewServeMux()
	mainRouter.Handle("/api/", http.StripPrefix("/api", subrouter))

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	chirpStore := chirp.NewStore(s.db)
	chirpHandler := chirp.NewHandler(chirpStore)
	chirpHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, mainRouter)
}
