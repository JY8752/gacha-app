package registory

import (
	user_repository "JY8752/gacha-app/domain/repository/user"
	useritem_repository "JY8752/gacha-app/domain/repository/useritem"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	user_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/user"
	useritem_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/useritem"
)

type UserServiceRegistory interface {
	User() user_repository.UserRepository
	UserItem() useritem_repository.UserItemRepository
}

type userServiceRegistory struct {
	client *datastore.MongoClient
}

func NewUserServiceRegistory(mongo *datastore.MongoClient) UserServiceRegistory {
	return &userServiceRegistory{mongo}
}

func (u *userServiceRegistory) User() user_repository.UserRepository {
	return user_datastore.NewUserRepository(u.client)
}

func (u *userServiceRegistory) UserItem() useritem_repository.UserItemRepository {
	return useritem_datastore.NewUserItemRepository(u.client)
}
