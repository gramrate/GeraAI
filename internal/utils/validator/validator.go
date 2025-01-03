package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate = validator.New()

// ValidateStruct validates a struct using the `validator` package
func ValidateStruct(data interface{}) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}
	return errors
}

// ValidateInterests проверяет, является ли Interests строкой в формате "[...;...]"
func ValidateInterests(interests string) error {
	// Проверяем, начинается ли строка с "[" и заканчивается на "]"
	if !strings.HasPrefix(interests, "[") || !strings.HasSuffix(interests, "]") {
		return errors.New("interests must be enclosed in square brackets")
	}

	// Убираем квадратные скобки
	content := interests[1 : len(interests)-1]

	// Разбиваем строку по разделителю ";"
	elements := strings.Split(content, ";")

	// Проверяем, что в массиве есть хотя бы один элемент
	if len(elements) == 0 {
		return errors.New("interests cannot be empty")
	}
	return nil
}
