package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func MobileValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^1[3-9]\d{9}$`, mobile)
	if !ok {
		return false
	}
	return true
}
