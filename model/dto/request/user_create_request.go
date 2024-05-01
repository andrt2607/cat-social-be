package requestdto

type UserCreateRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
