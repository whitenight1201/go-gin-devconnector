package repository

import (
	"errors"

	"github.com/whitenight1201/go-devconnector/pkg/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(profile entity.Profile) (entity.Profile, error)
	Update(profile entity.Profile) (entity.Profile, error)
	FindById(userId string) (entity.Profile, error)
	GetAll(userId string) ([]entity.Profile, error)
	Delete(profileId string, userId string) (entity.Profile, error)
}

type ProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &ProfileRepositoryImpl{
		db: db,
	}
}

func (profleRepo *ProfileRepositoryImpl) Create(profile entity.Profile) (entity.Profile, error) {
	//profleRepo.db.LogMode(true)
	if err := profleRepo.db.Debug().Create(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}

func (profleRepo *ProfileRepositoryImpl) Update(profile entity.Profile) (entity.Profile, error) {
	result := profleRepo.db.Where("id = ? AND user_id = ?", profile.ID, profile.UserID).Updates(&profile)

	if result.RowsAffected == 0 {
		return profile, errors.New("you dont have access to update this profile")
	}

	profleRepo.db.Preload("User").Find(&profile)
	return profile, nil
}

func (profleRepo *ProfileRepositoryImpl) FindById(userId string) (entity.Profile, error) {
	var profile entity.Profile
	result := profleRepo.db.Where("user_id = ?", userId).First(&profile)

	if result.RowsAffected == 0 {
		return profile, result.Error
	}

	return profile, nil
}

func (profleRepo *ProfileRepositoryImpl) GetAll(userId string) ([]entity.Profile, error) {
	var profile []entity.Profile
	if err := profleRepo.db.Where("user_id = ?", userId).Find(&profile).Error; err != nil {
		return profile, err
	}

	profleRepo.db.Preload("User").Find(&profile)
	return profile, nil
}

func (profleRepo *ProfileRepositoryImpl) Delete(profileId string, userId string) (entity.Profile, error) {
	var profile entity.Profile
	result := profleRepo.db.Where("id = ? AND user_id = ?", profileId, userId).Delete(&profile)
	if result.RowsAffected == 0 {
		return profile, errors.New("you dont have access to delete this profile")
	}

	profleRepo.db.Preload("User").Find(&profile)
	return profile, nil
}
