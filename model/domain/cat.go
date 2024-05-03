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
	OwnerId     string
	isMatched   bool
	isDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
