package domain

import (
	"time"
)

type Match struct {
	Id          	int
	OwnerId      	string
	CatId        	string
	LikedOwnerId 	string
	LikedCatId  	string
	ApprovalStatus 	string
	Messaged   		string
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
}
