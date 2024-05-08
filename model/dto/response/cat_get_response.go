package responsedto

type CatGetResponse struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	AgeInMonth  int      `json:"ageInMonth"`
	ImageUrls   []string `json:"imageUrls"`
	Description string   `json:"description"`
	HasMatched  bool     `json:"hasMatched"`
	CreatedAt   string   `json:"createdAt"`
}
