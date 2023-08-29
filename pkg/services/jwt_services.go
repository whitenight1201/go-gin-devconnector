package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/whitenight1201/go-devconnector/pkg/exception"
)

type JWTServices interface {
	GenerateToken(userID string) string
	ValidateToken(tokens string, c *gin.Context) *jwt.Token
}

type JWTServicesImpl struct {
	issuer    string
	secretKey string
}

type jwtCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func NewJWTServices() JWTServices {
	return &JWTServicesImpl{
		issuer:    "admin",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "adminJWT"
	}
	return key
}

func (jwtServices *JWTServicesImpl) GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 2).Unix(),
			Issuer:    jwtServices.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokens.SignedString([]byte(jwtServices.secretKey))
	exception.PanicIfNeeded(err)

	return token
}

func (jwtServices *JWTServicesImpl) ValidateToken(tokens string, c *gin.Context) *jwt.Token {
	t, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(jwtServices.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t
}
