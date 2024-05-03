package requestdto

type MatchApproveRequest struct {
	MatchId string `validate:"required" json:"matchId"`
}
