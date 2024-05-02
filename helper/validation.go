package helper

import "github.com/go-playground/validator/v10"

// Validator adalah instance dari validator
var Validator *validator.Validate

func init() {
	Validator = validator.New()
}

// ValidateStruct digunakan untuk memvalidasi struct
func ValidateStruct(s interface{}) error {
	return Validator.Struct(s)
}
