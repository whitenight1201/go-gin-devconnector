package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/services"

	"log"
	"net/http"
)

func AuthorizeJWT(jwtServices services.JWTServices) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := response.BuildErrorResponse("Failed to process request", "No token provided")
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		token := jwtServices.ValidateToken(authHeader, c)
		if token != nil {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			res := response.BuildErrorResponse("Error", "Your token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
