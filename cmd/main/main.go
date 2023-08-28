package main

import (
	"github.com/whitenight1201/go-devconnector/pkg/database"
	"github.com/whitenight1201/go-devconnector/pkg/server"
)

// var sugarLogger *zap.SugaredLogger

func main() {
	// InitLogger()
	// defer sugarLogger.Sync()

	database.DatabaseConnection()

	server.Start()
}

// func InitLogger() {
// 	logger, _ := zap.NewProduction()
// 	sugarLogger = logger.Sugar()
// }
