package model

import "time"

type Gacha struct {
	Id      string
	GachaId string
	Name    string
	Items   []struct {
		ItemId string
		Weight int
	}
	UpdatedAt time.Time
	CreatedAt time.Time
}
