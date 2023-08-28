package server

import (
	conf "github.com/whitenight1201/go-devconnector/pkg/config"
	"github.com/whitenight1201/go-devconnector/pkg/database"
)

func Start(cfg conf.Config) {
	jwtSetup(cfg)

	database.DatabaseConnection(cfg)

	router := setRouter()

	// Start listening and serving requests
	router.Run("localhost:5000")
}
