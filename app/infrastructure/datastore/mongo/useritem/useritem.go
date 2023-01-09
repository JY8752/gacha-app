package datastore

import (
	"JY8752/gacha-app/constant"
	model "JY8752/gacha-app/domain/model/useritem"
	repository "JY8752/gacha-app/domain/repository/useritem"
	applicationerror "JY8752/gacha-app/error"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "UserItems"

type UserItem struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"usrId"`
	ItemId    string             `bson:"itmId"`
	Count     int                `bson:"cnt"`
	UpdatedAt time.Time          `bson:"updAt"`
	CreatedAt time.Time          `bson:"crtAt"`
}

type userItemRepository struct {
	client *datastore.MongoClient
}

func NewUserItemRepository(c *datastore.MongoClient) repository.UserItemRepository {
	return &userItemRepository{client: c}
}

func (u *userItemRepository) Create(ctx context.Context, userId primitive.ObjectID, itemId string, count int, time time.Time) (primitive.ObjectID, error) {
	doc := UserItem{
		Id:        primitive.NewObjectID(),
		UserId:    userId,
		ItemId:    itemId,
		Count:     1,
		UpdatedAt: time,
		CreatedAt: time,
	}

	result, err := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).InsertOne(ctx, &doc)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, applicationerror.NewApplicationError(fmt.Sprintf("Fail cast to ObjectId. result: %v\n", result), nil)
}

func (u *userItemRepository) IncrementCount(ctx context.Context, userId primitive.ObjectID, itemId string, time time.Time) error {
	filter := bson.D{{Key: "usrId", Value: userId}, {Key: "itmId", Value: itemId}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "cnt", Value: 1}}}, {Key: "$set", Value: bson.D{{Key: "updAt", Value: time}}}}
	result := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOneAndUpdate(ctx, filter, update)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return applicationerror.NewApplicationError("", result.Err())
		}
		log.Fatal(result.Err())
	}
	return nil
}

func (u *userItemRepository) List(ctx context.Context, userId primitive.ObjectID) []model.UserItem {
	filter := bson.D{{Key: "usrId", Value: userId}}
	cursol, err := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []UserItem
	if err = cursol.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	var models []model.UserItem
	for _, result := range results {
		models = append(models, *toModel(result))
	}

	return models
}

func toModel(doc UserItem) *model.UserItem {
	return &model.UserItem{
		Id:        doc.Id.Hex(),
		UserId:    doc.UserId.Hex(),
		ItemId:    doc.ItemId,
		Count:     doc.Count,
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

// for test

func (u *userItemRepository) FindById(ctx context.Context, id primitive.ObjectID) (*model.UserItem, error) {
	result := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, bson.D{{Key: "_id", Value: id}})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, applicationerror.NewApplicationError("", result.Err())
		}
		log.Fatal(result.Err())
	}

	var userItem UserItem
	if err := result.Decode(&userItem); err != nil {
		log.Fatal(err)
	}

	return toModel(userItem), nil
}

func (u *userItemRepository) FindByUserIdAndItemId(ctx context.Context, userId primitive.ObjectID, itemId string) (*model.UserItem, error) {
	filter := bson.D{{Key: "usrId", Value: userId}, {Key: "itmId", Value: itemId}}
	result := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, applicationerror.NewApplicationError("", result.Err())
		}
		log.Fatal(result.Err())
	}

	var userItem UserItem
	if err := result.Decode(&userItem); err != nil {
		log.Fatal(err)
	}

	return toModel(userItem), nil
}

func (u *userItemRepository) Delete(ctx context.Context) {
	if _, err := u.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).DeleteMany(ctx, bson.D{{}}); err != nil {
		log.Fatal(err)
	}
}
