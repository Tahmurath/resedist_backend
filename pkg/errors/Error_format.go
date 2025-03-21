package errors

import (
	"errors"
	"resedist/internal/providers/validation"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorFormat struct {
	ErrorsList map[string]string
}

func New() *ErrorFormat {
	return &ErrorFormat{
		ErrorsList: make(map[string]string),
	}
}

func (e *ErrorFormat) SetFromError(err error) map[string]string {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			e.Add(fieldError.Field(), e.GetErrorMsg(fieldError.Tag()))
		}
	}
	return e.Get()
	// fmt.Println(e.ErrorsList) // فقط برای دیباگ
}

func (e *ErrorFormat) Add(key string, value string) {
	if e.ErrorsList == nil {
		e.ErrorsList = make(map[string]string)
	}
	e.ErrorsList[strings.ToLower(key)] = value
}

func (e *ErrorFormat) GetErrorMsg(tag string) string {
	if msg, ok := validation.ErrorMessages()[tag]; ok {
		return msg
	}
	return "unknown validation error"
}

func (e *ErrorFormat) Get() map[string]string {
	return e.ErrorsList
}
