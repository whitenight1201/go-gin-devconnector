package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
	"github.com/whitenight1201/go-devconnector/pkg/middleware"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"
)

type ProfileController interface {
	ProfileRouters(group *gin.Engine)
	CreateProfileUser(c *gin.Context)
}

type ProfielControllerImpl struct {
	profileServices services.ProfileServices
	jwtServices     services.JWTServices
}

func NewProfileController(profileServices services.ProfileServices, jwtServices services.JWTServices) ProfileController {
	return &ProfielControllerImpl{
		profileServices: profileServices,
		jwtServices:     jwtServices,
	}
}

func (profileController *ProfielControllerImpl) ProfileRouters(router *gin.Engine) {
	route := router.Group("/", middleware.AuthorizeJWT(profileController.jwtServices))
	route.POST("/", profileController.CreateProfileUser)
}

func (profileController *ProfielControllerImpl) CreateProfileUser(c *gin.Context) {
	var profileRequest dto.CreateProfileRequest

	if err := c.ShouldBind(&profileRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	claims := profileController.jwtServices.GetClaimsJWT(c)
	id := fmt.Sprintf("%v", claims["user_id"])
	_id, _ := strconv.ParseInt(id, 0, 64)

	profileRequest.UserID = _id
	result, err := profileController.profileServices.CreateProfile(profileRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create profile", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
