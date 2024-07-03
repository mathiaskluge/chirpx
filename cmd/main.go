package main

import (
	"log"

	"github.com/mathiaskluge/chirpx/cmd/api"
	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/db"
)

func main() {

	db, err := db.NewDB(config.Env.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	srv := api.NewAPIServer(config.Env.Port, db)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
