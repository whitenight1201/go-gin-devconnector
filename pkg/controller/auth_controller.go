package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"
)

type AuthController interface {
	AuthRoutes(group *gin.RouterGroup)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	authServices services.AuthServices
}

func NewAuthController(authServices services.AuthServices) AuthController {
	return &AuthControllerImpl{
		authServices: authServices,
	}
}

func (authController *AuthControllerImpl) AuthRoutes(group *gin.RouterGroup) {
	route := group.Group("/")
	route.POST("/auth", authController.Login)
	route.POST("/users", authController.Register)
}

func (authController *AuthControllerImpl) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	if err := c.ShouldBind(&registerRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := authController.authServices.RegisterUser(registerRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusCreated, res)
}

func (authController *AuthControllerImpl) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err := authController.authServices.VerifyCredential(loginRequest)
	if err != nil {
		res := response.BuildErrorResponse("Failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	res := response.BuildSuccessResponse("Success", user)
	c.JSON(http.StatusOK, res)
}
