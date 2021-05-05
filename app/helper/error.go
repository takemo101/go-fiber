package helper

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func ErrorsToMap(errors error) map[string]string {
	result := make(map[string]string)
	for key, err := range errors.(validation.Errors) {
		result[key] = err.Error()
	}
	return result
}
