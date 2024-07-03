package main

import (
	"log"

	"github.com/mathiaskluge/chirpx/cmd/api"
)

func main() {

	srv := api.NewAPIServer(":8080")

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
