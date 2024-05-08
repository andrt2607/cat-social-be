package domain

import (
	"time"
)

type Cat struct {
	Id          string
	Name        string
	Race        string
	Sex         string
	AgeInMonth  string
	Description string
	ImageUrls   string
	OwnerId     string
	HasMatched  bool
	IsDeleted   bool
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time
}

type CatCheck struct {
	MatchCat Cat
	UserCat  Cat
	Response string
}
