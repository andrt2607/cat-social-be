package domain

import (
	"time"
)

type Cat struct {
	Id          int
	Name        string
	Race        string
	Sex         string
	BirthDate   string
	Description string
	ImageUrl    string
	OwnerId     string
	isMatched   bool
	isDeleted   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
