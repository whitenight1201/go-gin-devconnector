package services

import (
	"github.com/mashingan/smapping"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
	"github.com/whitenight1201/go-devconnector/pkg/entity"
	"github.com/whitenight1201/go-devconnector/pkg/repository"
	"github.com/whitenight1201/go-devconnector/pkg/response"
	"github.com/whitenight1201/go-devconnector/pkg/validations"
)

type ProfileServices interface {
	CreateProfile(profileRequest dto.CreateProfileRequest) (*response.ProfileResponse, error)
	UpdateProfile(profileRequest dto.UpdateProfileRequest) (*response.ProfileResponse, error)
	FindProfileById(profileId string, userId string) (*response.ProfileResponse, error)
	GetAllProfile(userId string) (*[]response.ProfileResponse, error)
	DeleteById(profileId string, userId string) (*response.ProfileResponse, error)
}

type ProfileServicesImpl struct {
	profileRepo repository.ProfileRepository
}

func NewProfileServices(profileRepository repository.ProfileRepository) ProfileServices {
	return &ProfileServicesImpl{
		profileRepo: profileRepository,
	}
}

func (profileServices *ProfileServicesImpl) CreateProfile(profileRequest dto.CreateProfileRequest) (*response.ProfileResponse, error) {
	var profile entity.Profile
	if err := validations.ValidateCreateProfile(profileRequest); err != nil {
		return nil, err
	}

	if err := smapping.FillStruct(&profile, smapping.MapFields(&profileRequest)); err != nil {
		return nil, err
	}

	result, err := profileServices.profileRepo.Create(profile)
	if err != nil {
		return nil, err
	}

	res := response.NewProfileResponse(result)
	return &res, nil
}

func (profileServices *ProfileServicesImpl) UpdateProfile(profileRequest dto.UpdateProfileRequest) (*response.ProfileResponse, error) {
	var profile entity.Profile
	if err := smapping.FillStruct(&profile, smapping.MapFields(&profileRequest)); err != nil {
		return nil, err
	}

	result, err := profileServices.profileRepo.Update(profile)
	if err != nil {
		return nil, err
	}

	res := response.NewProfileResponse(result)
	return &res, nil
}

func (profileServices *ProfileServicesImpl) FindProfileById(profileId string, userId string) (*response.ProfileResponse, error) {
	result, err := profileServices.profileRepo.FindById(profileId, userId)
	if err != nil {
		return nil, err
	}

	res := response.NewProfileResponse(result)
	return &res, nil
}

func (profileServices *ProfileServicesImpl) GetAllProfile(userId string) (*[]response.ProfileResponse, error) {
	result, err := profileServices.profileRepo.GetAll(userId)
	if err != nil {
		return nil, err
	}

	res := response.NewProfileResponseArray(result)
	return &res, nil
}

func (profileServices *ProfileServicesImpl) DeleteById(profileId string, userId string) (*response.ProfileResponse, error) {
	result, err := profileServices.profileRepo.Delete(profileId, userId)
	if err != nil {
		return nil, err
	}

	res := response.NewProfileResponse(result)
	return &res, nil
}
