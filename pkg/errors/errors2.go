package errors

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"resedist/internal/providers/validation"
	"strings"
)

type Error2 struct {
	ErrorsList map[string]string
}

func New() *Error2 {
	return &Error2{
		ErrorsList: make(map[string]string),
	}
}

func (e *Error2) SetFromError(err error) map[string]string {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			e.Add(fieldError.Field(), e.GetErrorMsg(fieldError.Tag()))
		}
	}
	return e.Get()
	// fmt.Println(e.ErrorsList) // فقط برای دیباگ
}

func (e *Error2) Add(key string, value string) {
	if e.ErrorsList == nil {
		e.ErrorsList = make(map[string]string)
	}
	e.ErrorsList[strings.ToLower(key)] = value
}

func (e *Error2) GetErrorMsg(tag string) string {
	if msg, ok := validation.ErrorMessages()[tag]; ok {
		return msg
	}
	return "unknown validation error"
}

func (e *Error2) Get() map[string]string {
	return e.ErrorsList
}
