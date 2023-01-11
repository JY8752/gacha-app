package model

import (
	model "JY8752/gacha-app/domain/model/item"
	"JY8752/gacha-app/util"
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

func (g *Gacha) Lottery() string {
	var itemIds []string
	var weights []int
	for _, i := range g.Items {
		itemIds = append(itemIds, i.ItemId)
		weights = append(weights, i.Weight)
	}

	index := util.BinarySearchLottery(weights)

	return itemIds[index]
}
