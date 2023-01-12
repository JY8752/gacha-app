package controller

import (
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	"JY8752/gacha-app/infrastructure/datastore/redis"
	"JY8752/gacha-app/pkg/grpc/gacha"
	"JY8752/gacha-app/pkg/grpc/user"
	"JY8752/gacha-app/registory"

	"google.golang.org/grpc"
)

func RegisterController(s grpc.ServiceRegistrar, mongo *datastore.MongoClient, redis *redis.RedisClient) {
	registory := registory.NewServiceRegistory(mongo, redis)

	user.RegisterUserServer(s, NewUserController(registory))
	gacha.RegisterGachaServer(s, NewGachaController(registory))
}
