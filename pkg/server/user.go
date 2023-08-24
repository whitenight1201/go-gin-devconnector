package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/controllers"
	"github.com/whitenight1201/go-devconnector/pkg/models"
)

func signUp(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	res, err := controllers.AddUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
	return
}
