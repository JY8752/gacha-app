package controller

import (
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	"JY8752/gacha-app/pkg/grpc/user"
	"JY8752/gacha-app/registory"

	"google.golang.org/grpc"
)

func RegisterController(s grpc.ServiceRegistrar, mongo *datastore.MongoClient) {
	userRegistory := registory.NewUserServiceRegistory(mongo)

	user.RegisterUserServer(s, NewUserController(userRegistory.User()))
}
