package repository

import (
	model "JY8752/gacha-app/domain/model/useritem"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserItemRepository interface {
	Create(ctx context.Context, userId primitive.ObjectID, itemId string, count int, time time.Time) (primitive.ObjectID, error)
	IncrementCount(ctx context.Context, userId primitive.ObjectID, itemId string, time time.Time) error
	List(ctx context.Context, userId primitive.ObjectID) []model.UserItem
	FindById(ctx context.Context, id primitive.ObjectID) (*model.UserItem, error)
	Delete(ctx context.Context)
	FindByUserIdAndItemId(ctx context.Context, userId primitive.ObjectID, itemId string) (*model.UserItem, error)
}
