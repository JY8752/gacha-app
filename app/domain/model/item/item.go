package model

import "time"

type Item struct {
	Id        string
	ItemId    string
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type ItemWith struct {
	ItemId string
	Weight int
}
