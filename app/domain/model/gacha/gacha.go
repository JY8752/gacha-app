package model

import (
	model "JY8752/gacha-app/domain/model/item"
	"time"
)

type Gacha struct {
	Id        string
	GachaId   string
	Name      string
	Items     []model.ItemWith
	UpdatedAt time.Time
	CreatedAt time.Time
}
