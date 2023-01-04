package repository

import (
	model "JY8752/gacha-app/domain/model/item"
	"context"
	"time"
)

type ItemRepository interface {
	Create(ctx context.Context, itemId, name string, time time.Time) (string, error)
	FindById(ctx context.Context, id string) (*model.Item, error)
}
