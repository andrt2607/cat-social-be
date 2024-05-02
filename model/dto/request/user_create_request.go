package requestdto

import "github.com/go-playground/validator/v10"

type UserCreateRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Name     string `validate:"required,min=5,max=50" json:"name"`
	Password string `validate:"required,min=5,max=15" json:"password"`
}

func (r *UserCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
