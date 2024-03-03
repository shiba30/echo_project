package validator

import (
	"regexp"
	"twitter_echo_backend/model"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

// コンストラクタ
func NewUserValidator() UserValidator {
	return &userValidator{}
}

func (tv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.DisplayName,
			validation.Required.Error("DisplayName is required"),
			validation.Length(1, 50).Error("name must be between 1 and 50 characters long"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("パスワードは必須です"),
			validation.Length(8, 100).Error("パスワードは8文字以上である必要があります"),
			is.Alphanumeric.Error("パスワードには半角英数字を含める必要があります"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("パスワードには大文字を含める必要があります"),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("パスワードには小文字を含める必要があります"),
			validation.Match(regexp.MustCompile(`[!?-_@]`)).Error("パスワードには以下の記号のいずれかを含める必要があります: !?-_@"),
		),
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("invalid Email format"),
		),
	)
}
