package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"time"

	"github.com/cristalhq/jwt/v3"
	conf "github.com/whitenight1201/go-devconnector/pkg/config"
	"github.com/whitenight1201/go-devconnector/pkg/models"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

// Function jwtSetup() will only create signer and verifier that will later be used in authentication.
// we can call this function from internal/server/server/go when starting server
func jwtSetup(cfg conf.Config) {
	var err error
	key := []byte(cfg.JwtSecret)

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Panic("Error creating JWT signer")
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Panic("Error creating JWT verifier")
	}
}

// This function generate tokens
func generateJWT(user *models.User) string {
	claims := &jwt.RegisteredClaims{
		ID:        fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Panic("Error building JWT")
	}
	return token.String()
}

// This function will receive token in the form of string, verify its signature, extract ID from claims and if everything is ok, user’s ID will be returned as int:
func verifyJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse([]byte(tokenStr))
	if err != nil {
		log.Fatalf("Error parsing JWT")
		return 0, err
	}
	if err := jwtVerifier.Verify(token.Payload(), token.Signature()); err != nil {
		log.Fatalf("Error verifying token")
		return 0, err
	}
	var claims jwt.StandardClaims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		log.Fatalf("Error unmarshalling JWT claims")
		return 0, err
	}
	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("token expired")
	}
	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Fatalf("Error converting claims ID to number")
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}
