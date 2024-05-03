package responsedto

import "time"

type MatchApproveResponse struct {
	Id			int       `json:"id"`
	CatId		string	  `json:"catId"`
	LikedCatId	string	  `json:"likedCatId"`
	CreatedAt	time.Time `json:"createdAt"`
	UpdatedAt	time.Time `json:"updatedAt"`
}