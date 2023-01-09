package service

import (
	model "JY8752/gacha-app/domain/model/user"
	repository "JY8752/gacha-app/domain/repository/user"
	"context"
	"time"
)

type UserService interface {
	Create(ctx context.Context, name string, time time.Time) (*model.User, error)
}

type userService struct {
	userRep repository.UserRepository
}

func NewUserService(rep repository.UserRepository) UserService {
	return &userService{rep}
}

func (u *userService) Create(ctx context.Context, name string, time time.Time) (*model.User, error) {
	oid, err := u.userRep.Create(ctx, name, time)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: oid, Name: name, UpdatedAt: time, CreatedAt: time}, nil
}

func (u *userService) ListUserItems() {

}
