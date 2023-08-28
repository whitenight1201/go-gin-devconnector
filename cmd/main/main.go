package main

import (
	conf "github.com/whitenight1201/go-devconnector/pkg/config"
	"github.com/whitenight1201/go-devconnector/pkg/server"
)

// var sugarLogger *zap.SugaredLogger

func main() {
	// InitLogger()
	// defer sugarLogger.Sync()

	server.Start(conf.NewConfig())
}

// func InitLogger() {
// 	logger, _ := zap.NewProduction()
// 	sugarLogger = logger.Sugar()
// }
