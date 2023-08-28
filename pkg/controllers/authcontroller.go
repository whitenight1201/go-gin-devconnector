package controllers

import (
	"crypto/rand"
	"errors"

	"github.com/whitenight1201/go-devconnector/pkg/database"
	"github.com/whitenight1201/go-devconnector/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AddUser(user *models.User) error {
	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	toHash := append([]byte(user.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Salt = salt
	user.HashedPassword = hashedPassword

	var result models.User
	dbResult := database.DB.Where("useremail = ?", user.Useremail).First(&result)
	//fmt.Println(errors.Is(dbResult.Error, gorm.ErrRecordNotFound))

	if dbResult.Error == nil {
		return errors.New("user already exist")
	}

	if dbResult.Error.Error() == gorm.ErrRecordNotFound.Error() {
		res := database.DB.Create(user)
		if res.RowsAffected == 0 {
			return errors.New("errors creating user")
		}
	}

	return nil
}

func Authenticate(useremail, password string) (*models.User, error) {
	var user models.User
	res := database.DB.Where("useremail = ?", useremail).First(&user)
	if res.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	salted := append([]byte(password), user.Salt...)
	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, salted); err != nil {
		return nil, err
	}

	return &user, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
