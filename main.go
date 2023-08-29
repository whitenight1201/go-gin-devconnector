package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/config"
)

func main() {
	// Create default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	config.DatabaseConnection()

	// Create API route group
	api := router.Group("/api")
	{
		// Add /hello GET route to router and define route handler function
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	//Start listening and serving requests
	router.Run("localhost:5000")
}
