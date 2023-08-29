package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"
)

type AuthController interface {
	AuthRoutes(group *gin.Engine)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	authServices services.AuthServices
	jwtServices  services.JWTServices
}

func NewAuthController(authServices services.AuthServices, jwtServices services.JWTServices) AuthController {
	return &AuthControllerImpl{
		authServices: authServices,
		jwtServices:  jwtServices,
	}
}

func (authController *AuthControllerImpl) AuthRoutes(router *gin.Engine) {
	route := router.Group("/api")
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

	user, err := authController.authServices.RegisterUser(registerRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	//Generate Token
	token := authController.jwtServices.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token

	res := response.BuildSuccessResponse("Success", user)
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

	//Generate Token
	token := authController.jwtServices.GenerateToken(strconv.FormatInt(user.ID, 10))
	user.Token = token

	res := response.BuildSuccessResponse("Success", user)
	c.JSON(http.StatusOK, res)
}
