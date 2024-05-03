package requestdto

import "github.com/go-playground/validator/v10"

type MatchCreateRequest struct {
	matchId    string `validate:"required" json:"match_approve_id"`
}