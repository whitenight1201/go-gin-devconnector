package services

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mashingan/smapping"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
	"github.com/whitenight1201/go-devconnector/pkg/repository"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/validations"

	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthServices interface {
	RegisterUser(registerRequest dto.RegisterRequest) (*response.UserResponse, error)
	VerifyCredential(loginRequest dto.LoginRequest) (*response.UserResponse, error)
}

type AuthServicesImpl struct {
	userRepo repository.UserRepository
}

func NewAuthServices(userRepository repository.UserRepository) AuthServices {
	return &AuthServicesImpl{
		userRepo: userRepository,
	}
}
func (authServices *AuthServicesImpl) RegisterUser(registerRequest dto.RegisterRequest) (*response.UserResponse, error) {
	if errValidate := validations.ValidateRegister(registerRequest); errValidate != nil {
		return nil, errValidate
	}

	user, errFind := authServices.userRepo.FindByEmail(registerRequest.Email)
	if errFind == nil {
		return nil, errors.New(fmt.Sprintf("User %v already exists", user.Name))
	}

	if errFind != nil && !errors.Is(errFind, gorm.ErrRecordNotFound) {
		return nil, errFind
	}

	if errMap := smapping.FillStruct(&user, smapping.MapFields(&registerRequest)); errMap != nil {
		log.Fatal("failed mapping %v", errMap)
		return nil, errMap
	}

	result, err := authServices.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	res := response.NewUserResponse(result)
	return &res, nil
}

func (authServices *AuthServicesImpl) VerifyCredential(loginRequest dto.LoginRequest) (*response.UserResponse, error) {
	if errValidate := validations.ValidateLogin(loginRequest); errValidate != nil {
		return nil, errValidate
	}

	user, err := authServices.userRepo.FindByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	isValidPass := comparePassword(user.HashedPassword, []byte(loginRequest.Password))
	if !isValidPass {
		return nil, errors.New("wrong username and password")
	}

	res := response.NewUserResponse(user)
	return &res, nil
}

func comparePassword(hashPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPass, plainPass)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
