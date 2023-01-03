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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "Users"

type user struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"nm"`
	UpdatedAt time.Time          `bson:"updAt"`
	CreatedAt time.Time          `bson:"crtAt"`
}

type userRepository struct {
	client *datastore.MongoClient
}

func NewUserRepository(client *datastore.MongoClient) repository.UserRepository {
	return &userRepository{client: client}
}

func (u *userRepository) Create(ctx context.Context, name string, time time.Time) (string, error) {
	doc := &user{Id: primitive.NewObjectID(), Name: name, UpdatedAt: time, CreatedAt: time}
	result, err := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).InsertOne(ctx, doc)

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

	filter := bson.D{{Key: "_id", Value: oid}}
	var user user
	if err := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, applicationerror.NewApplicationError(fmt.Sprintf("Fail findById id: %s\n", id), err)
	}

	return &model.User{
		Id:        user.Id.Hex(),
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}, nil
}
