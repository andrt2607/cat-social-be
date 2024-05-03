package responsedto

import "time"

type MatchCreateResponse struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
