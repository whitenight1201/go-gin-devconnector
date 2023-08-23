package main

import (
	"github.com/whitenight1201/go-devconnector/pkg/database"
	"github.com/whitenight1201/go-devconnector/pkg/server"
)

func main() {
	database.DatabaseConnection()

	server.Start()
}
