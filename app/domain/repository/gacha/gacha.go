package repository

import (
	gacha_model "JY8752/gacha-app/domain/model/gacha"
	item_model "JY8752/gacha-app/domain/model/item"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GachaRepository interface {
	FindByGachaId(ctx context.Context, gachaId string) (*gacha_model.Gacha, error)
	Create(ctx context.Context, gachaId, name string, items []item_model.ItemWith, time time.Time) (primitive.ObjectID, error)
}

type GachaHistoryRepository interface {
	Add(ctx context.Context, gachaId, itemId string, userId primitive.ObjectID) error
	Get(ctx context.Context, gachaId string, userId primitive.ObjectID) string
}
