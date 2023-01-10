package registory

import (
	service "JY8752/gacha-app/domain/service/user"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	user_repository "JY8752/gacha-app/infrastructure/datastore/mongo/user"
	useritem_repository "JY8752/gacha-app/infrastructure/datastore/mongo/useritem"
)

type UserServiceRegistory interface {
	User() *service.UserService
}

type userServiceRegistory struct {
	client *datastore.MongoClient
}

func NewUserServiceRegistory(mongo *datastore.MongoClient) userServiceRegistory {
	return userServiceRegistory{mongo}
}

func (u *userServiceRegistory) User() service.UserService {
	ur := user_repository.NewUserRepository(u.client)
	uir := useritem_repository.NewUserItemRepository(u.client)
	return service.NewUserService(ur, uir)
}
