package service

import (
	model "JY8752/gacha-app/domain/model/item"
	repository "JY8752/gacha-app/domain/repository/item"
	"JY8752/gacha-app/registory"
	"context"
)

type ItemService interface {
	FindInItemId(ctx context.Context, ids []string) map[string]model.Item
}

type itemService struct {
	itemRep repository.ItemRepository
}

func NewItemService(r registory.ServiceRegistory) ItemService {
	return &itemService{r.Item()}
}

func (i *itemService) FindInItemId(ctx context.Context, ids []string) map[string]model.Item {
	items := i.itemRep.FindInItemId(ctx, ids)

	itemMap := make(map[string]model.Item)
	for _, item := range items {
		itemMap[item.ItemId] = item
	}
	return itemMap
}
