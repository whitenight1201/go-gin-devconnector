package controller

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/middleware"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"

	"net/http"
)

type UserController interface {
	UserRoutes(group *gin.Engine)
	FetchUser(c *gin.Context)
}

type UserControllerImpl struct {
	userServices services.UserServices
	jwtServices  services.JWTServices
}

func NewUserController(userServices services.UserServices, jwtServices services.JWTServices) UserController {
	return &UserControllerImpl{
		userServices: userServices,
		jwtServices:  jwtServices,
	}
}

func (userController *UserControllerImpl) UserRoutes(router *gin.Engine) {
	route := router.Group("/api", middleware.AuthorizeJWT(userController.jwtServices))
	route.GET("/auth", userController.FetchUser)
}

func (userController *UserControllerImpl) FetchUser(c *gin.Context) {
	header := c.GetHeader("x-auth-token")
	token := userController.jwtServices.ValidateToken(header, c)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	user, err := userController.userServices.FindUserById(id)
	if err != nil {
		res := response.BuildErrorResponse("Error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	res := response.BuildSuccessResponse("Success", user)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
