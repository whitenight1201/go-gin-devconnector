package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/config"
	"github.com/whitenight1201/go-devconnector/pkg/controller"
	"github.com/whitenight1201/go-devconnector/pkg/repository"
	"github.com/whitenight1201/go-devconnector/pkg/services"
)

func main() {
	db := config.DatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	userServices := services.NewUserServices(userRepository)
	jwtServices := services.NewJWTServices()
	authServices := services.NewAuthServices(userRepository)
	authController := controller.NewAuthController(authServices, jwtServices)
	userController := controller.NewUserController(userServices, jwtServices)

	// Create default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Add /hello GET route to router and define route handler function
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "hello world"})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	// Create API route group
	api := router.Group("/api")

	authController.AuthRoutes(api)
	userController.UserRoutes(api)
	//Start listening and serving requests
	router.Run("localhost:5000")
}
