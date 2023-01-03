package model

import "time"

type UserItem struct {
	Id        string
	UserId    string
	ItemId    string
	Count     int
	UpdatedAt time.Time
	CreatedAt time.Time
}
