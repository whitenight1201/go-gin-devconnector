package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/middleware"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"

	"net/http"
)

type UserController interface {
	UserRoutes(group *gin.Engine)
	CurrentUser(c *gin.Context)
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
	route.GET("/auth", userController.CurrentUser)
}

func (userController *UserControllerImpl) CurrentUser(c *gin.Context) {
	claims := userController.jwtServices.GetClaimsJWT(c)
	id := fmt.Sprintf("%v", claims["user_id"])

	user, err := userController.userServices.FindUserById(id)
	if err != nil {
		res := response.BuildErrorResponse("Error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	res := response.BuildSuccessResponse("Success", user)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
