package validations

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
)

func ValidateCreateProfile(profile dto.CreateProfileRequest) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Status, validation.Required, validation.Length(10, 0)),
		validation.Field(&profile.Skills, validation.Required),
		validation.Field(&profile.UserID, validation.Required))
}
