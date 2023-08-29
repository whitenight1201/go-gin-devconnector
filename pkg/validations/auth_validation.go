package validations

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/whitenight1201/go-devconnector/pkg/dto"
)

func ValidateRegister(register dto.RegisterRequest) error {
	return validation.ValidateStruct(&register,
		validation.Field(&register.Name, validation.Required),
		validation.Field(&register.Email, validation.Required, is.Email),
		validation.Field(&register.Password, validation.Required, validation.Length(6, 0)))
}

func ValidateLogin(login dto.LoginRequest) error {
	return validation.ValidateStruct(&login,
		validation.Field(&login.Email, validation.Required, is.Email),
		validation.Field(&login.Password, validation.Required, validation.Length(6, 0)))
}
