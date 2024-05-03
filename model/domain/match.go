package domain

import (
	"time"
)

type Match struct {
	Id          	int
	ownerEmail      string
	catId        	string
	likedOwnerEmail string
	likedCatId  	string
	isApproved 		string
	Messaged   		string
	CreatedAt   	time.Time
	UpdatedAt   	time.Time
}
