package datastore

import (
	"JY8752/gacha-app/constant"
	model "JY8752/gacha-app/domain/model/user"
	repository "JY8752/gacha-app/domain/repository/user"
	applicationerror "JY8752/gacha-app/error"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "Users"

type user struct {
	id        primitive.ObjectID `bson:"_id"`
	name      string             `bson:"nm"`
	updatedAt time.Time          `bson:"updAt"`
	createdAt time.Time          `bson:"crtAt"`
}

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (u *userRepository) Create(ctx context.Context, name string, time time.Time) (string, error) {
	doc := &user{name: name, updatedAt: time, createdAt: time}
	result, err := datastore.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).InsertOne(ctx, doc)

	if err != nil {
		return "", applicationerror.NewApplicationError("Fail create user.", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", applicationerror.NewApplicationError(fmt.Sprintf("Fail cast to objectId. result: %v\n", result), nil)
}

func (u *userRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, applicationerror.NewApplicationError(fmt.Sprintf("argument id is not ObjectId. id: %s\n", id), err)
	}

	filter := &user{id: oid}
	m := &model.User{}
	if err := datastore.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, filter).Decode(m); err != nil {
		return nil, applicationerror.NewApplicationError(fmt.Sprintf("Fail findById id: %s\n", id), err)
	}
	return m, nil
}
