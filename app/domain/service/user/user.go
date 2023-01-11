package service

import (
	user "JY8752/gacha-app/domain/model/user"
	userItem "JY8752/gacha-app/domain/model/useritem"
	userRepository "JY8752/gacha-app/domain/repository/user"
	userItemRepository "JY8752/gacha-app/domain/repository/useritem"
	"JY8752/gacha-app/registory"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Create(ctx context.Context, name string, time time.Time) (*user.User, error)
	ListUserItems(ctx context.Context, userId primitive.ObjectID) []userItem.UserItem
}

type userService struct {
	userRep     userRepository.UserRepository
	userItemRep userItemRepository.UserItemRepository
}

func NewUserService(r registory.ServiceRegistory) UserService {
	return &userService{userRep: r.User(), userItemRep: r.UserItem()}
}

func (u *userService) Create(ctx context.Context, name string, time time.Time) (*user.User, error) {
	oid, err := u.userRep.Create(ctx, name, time)
	if err != nil {
		return nil, err
	}
	return &user.User{Id: oid, Name: name, UpdatedAt: time, CreatedAt: time}, nil
}

func (u *userService) ListUserItems(ctx context.Context, userId primitive.ObjectID) []userItem.UserItem {
	return u.userItemRep.List(ctx, userId)
}
