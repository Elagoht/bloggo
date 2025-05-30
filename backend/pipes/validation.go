package pipes

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func slugValidation(fl validator.FieldLevel) bool {
	return slugRegex.MatchString(fl.Field().String())
}

func createValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("slug", slugValidation)
	return validate
}

var defaultValidator = createValidator()

func GetValidator() *validator.Validate {
	return defaultValidator
}
