package main

import (
	"github.com/0xEg0x/api-students/api"
	"github.com/rs/zerolog/log"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("failed start server")
	}

}
