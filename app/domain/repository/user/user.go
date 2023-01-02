package repository

import (
	model "JY8752/gacha-app/domain/model/user"
	"context"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, name string, time time.Time) (string, error)
	FindById(ctx context.Context, id string) (*model.User, error)
}
