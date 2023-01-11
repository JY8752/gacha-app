package registory

import (
	item_repository "JY8752/gacha-app/domain/repository/item"
	user_repository "JY8752/gacha-app/domain/repository/user"
	useritem_repository "JY8752/gacha-app/domain/repository/useritem"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	item_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/item"
	user_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/user"
	useritem_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/useritem"
)

type ServiceRegistory interface {
	User() user_repository.UserRepository
	UserItem() useritem_repository.UserItemRepository
	Item() item_repository.ItemRepository
}

type serviceRegistory struct {
	client *datastore.MongoClient
}

func NewServiceRegistory(mongo *datastore.MongoClient) ServiceRegistory {
	return &serviceRegistory{mongo}
}

func (u *serviceRegistory) User() user_repository.UserRepository {
	return user_datastore.NewUserRepository(u.client)
}

func (u *serviceRegistory) UserItem() useritem_repository.UserItemRepository {
	return useritem_datastore.NewUserItemRepository(u.client)
}

func (u *serviceRegistory) Item() item_repository.ItemRepository {
	return item_datastore.NewItemRepository(u.client)
}
