package controllers

import (
	"crypto/rand"
	"errors"

	"github.com/whitenight1201/go-devconnector/pkg/database"
	"github.com/whitenight1201/go-devconnector/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(user *models.User) (*models.User, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return &models.User{}, err
	}
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return &models.User{}, err
	}
	user.Salt = salt
	user.HashedPassword = hashedPassword

	res := database.DB.Create(user)

	if res.RowsAffected == 0 {
		return &models.User{}, errors.New("errors creating user")
	}

	return user, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
