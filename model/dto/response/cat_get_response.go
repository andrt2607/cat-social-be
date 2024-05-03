package responsedto

import "time"

type CatGetResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   string    `json:"imageUrls"`
	Description string    `json:"description"`
	HasMatched  string    `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}
