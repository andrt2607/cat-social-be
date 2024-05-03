package requestdto

type MatchCreateRequest struct {
	MatchCatId    string `validate:"required" json:"matchCatId"`
	UserCatId     string `validate:"required" json:"userCatId"`
	Message 	  string `validate:"required,min=5,max=120" json:"message"`
}