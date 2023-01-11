package registory

import (
	gacha_repository "JY8752/gacha-app/domain/repository/gacha"
	item_repository "JY8752/gacha-app/domain/repository/item"
	user_repository "JY8752/gacha-app/domain/repository/user"
	useritem_repository "JY8752/gacha-app/domain/repository/useritem"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	gacha_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/gacha"
	item_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/item"
	user_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/user"
	useritem_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/useritem"
)

type ServiceRegistory interface {
	User() user_repository.UserRepository
	UserItem() useritem_repository.UserItemRepository
	Item() item_repository.ItemRepository
	Gacha() gacha_repository.GachaRepository
}

type serviceRegistory struct {
	client *datastore.MongoClient
}

func NewServiceRegistory(mongo *datastore.MongoClient) ServiceRegistory {
	return &serviceRegistory{mongo}
}

func (s *serviceRegistory) User() user_repository.UserRepository {
	return user_datastore.NewUserRepository(s.client)
}

func (s *serviceRegistory) UserItem() useritem_repository.UserItemRepository {
	return useritem_datastore.NewUserItemRepository(s.client)
}

func (s *serviceRegistory) Item() item_repository.ItemRepository {
	return item_datastore.NewItemRepository(s.client)
}

func (s *serviceRegistory) Gacha() gacha_repository.GachaRepository {
	return gacha_datastore.NewGachaRepository(s.client)
}
