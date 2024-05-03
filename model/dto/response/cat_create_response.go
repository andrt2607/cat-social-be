package responsedto

import "time"

type CatCreateResponse struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
