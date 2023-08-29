package services

import (
	"github.com/whitenight1201/go-devconnector/pkg/repository"
	"github.com/whitenight1201/go-devconnector/pkg/response"
)

type UserServices interface {
	FindUserById(id string) (*response.UserResponse, error)
}

type UserServicesImpl struct {
	userRepo repository.UserRepository
}

func NewUserServices(userRepository repository.UserRepository) UserServices {
	return &UserServicesImpl{
		userRepo: userRepository,
	}
}

func (userServices *UserServicesImpl) FindUserById(id string) (*response.UserResponse, error) {
	user, err := userServices.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	result := response.NewUserResponse(user)
	return &result, nil
}
