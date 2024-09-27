package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"resedist/internal/providers/validation"
	"strings"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromError(err error) {

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func Add(key string, value string) {

	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}
