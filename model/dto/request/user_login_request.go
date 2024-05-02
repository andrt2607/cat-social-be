package requestdto

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=5,max=15" json:"password"`
}
