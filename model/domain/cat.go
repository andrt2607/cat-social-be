package domain

import (
	"time"
)

type Cat struct {
	Id          int
	Name        string
	Race        string
	Sex         string
	AgeInMonth  string
	Description string
	ImageUrls   string
	OwnerId  	string
	IsMatched   bool
	IsDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CatCheck struct {
	MatchCat Cat
	UserCat Cat
	Response string
}