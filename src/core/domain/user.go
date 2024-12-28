package domain

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	custom_errors "github.com/hpaes/go-api-final-project/src/core/domain/customErrors"
	"github.com/hpaes/go-api-final-project/src/core/domain/vo"
)

type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name" validate:"required"`
	Email vo.Email `json:"email" validate:"required,email"`
	Age   int      `json:"age" validate:"required,gt=0"`
}

var (
	validate   *validator.Validate
	translator ut.Translator
)

func init() {
	validate = validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)
	translator, _ = uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, translator)
	registerCustomValidations()
	registerCustomTranslations()
}

func registerCustomValidations() {
	validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		email, ok := fl.Field().Interface().(vo.Email)
		if !ok {
			return false
		}
		return email.Validate() == nil
	})
}

func registerCustomTranslations() {
	registerTranslation("required", "{0} is a required field")
	registerTranslation("email", "{0} must be a valid email address")
	registerTranslation("gt", "{0} must be greater than {1}")
}

func registerTranslation(tag, message string) {
	validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field(), fe.Param())
		return t
	})
}

func newUser(id int, name, email string, age int) (*User, error) {
	emailVo, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	user := &User{
		Id:    id,
		Name:  name,
		Email: *emailVo,
		Age:   age,
	}

	err = validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, custom_errors.NewInvalidParameterError(err.Field(), err.Translate(translator))
		}
	}
	return user, nil
}

func Create(name, email string, age int) (*User, error) {
	user, err := newUser(0, name, email, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Restore(id int, name, email string, age int) (*User, error) {
	user, err := newUser(id, name, email, age)
	if err != nil {
		return nil, err
	}
	return user, nil
}
