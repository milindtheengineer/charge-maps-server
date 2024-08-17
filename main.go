package main

import (
	"github.com/milindtheengineer/charge-maps-server/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if err := config.InitialiseConfig(); err != nil {
		log.Panic().Msgf("Config could not be initialized due to %v", err)
	}

}
