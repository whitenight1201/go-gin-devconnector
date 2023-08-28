package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/controllers"
	"github.com/whitenight1201/go-devconnector/pkg/models"
)

func signUp(ctx *gin.Context) {
	user := new(models.User)

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err := controllers.AddUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": generateJWT(user),
	})
}

func signIn(ctx *gin.Context) {
	user := new(models.User)

	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	user, err := controllers.Authenticate(user.Useremail, user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sign in failed."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed in successfully.",
		"jwt": generateJWT(user),
	})
}
