package responsedto

import "time"

// MatchDetail adalah struktur untuk detail pencocokan kucing
type MatchDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   string    `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

// UserCatDetail adalah struktur untuk detail kucing pengguna
type UserCatDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	Description string    `json:"description"`
	AgeInMonth  int       `json:"ageInMonth"`
	ImageUrls   string    `json:"imageUrls"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}

// IssuedBy adalah struktur untuk penerbit informasi
type IssuedBy struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// MatchMessage adalah struktur untuk pesan pencocokan
type MatchGetResponse struct {
	ID             string        `json:"id"`
	IssuedBy       IssuedBy      `json:"issuedBy"`
	MatchCatDetail MatchDetail   `json:"matchCatDetail"`
	UserCatDetail  UserCatDetail `json:"userCatDetail"`
	Message        string        `json:"message"`
	CreatedAt      time.Time     `json:"createdAt"`
}
