package requestdto

import "github.com/go-playground/validator/v10"

type MatchCreateRequest struct {
	matchCatId    string `validate:"required" json:"match_cat_id"`
	userCatId     string `validate:"required" json:"user_cat_id"`
	message string `validate:"required,min=5,max=120" json:"message"`
}