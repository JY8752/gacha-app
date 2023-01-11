package datastore

import (
	"JY8752/gacha-app/constant"
	model "JY8752/gacha-app/domain/model/item"
	repository "JY8752/gacha-app/domain/repository/item"
	applicationerror "JY8752/gacha-app/error"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "Items"

type Item struct {
	Id        primitive.ObjectID `bson:"_id"`
	ItemId    string             `bson:"itmId"`
	Name      string             `bson:"nm"`
	UpdatedAt time.Time          `bson:"updAt"`
	CreatedAt time.Time          `bson:"crtAt"`
}

type itemRepository struct {
	client *datastore.MongoClient
}

func NewItemRepository(c *datastore.MongoClient) repository.ItemRepository {
	return &itemRepository{client: c}
}

func (ir *itemRepository) Create(ctx context.Context, itemId, name string, time time.Time) (string, error) {
	doc := Item{Id: primitive.NewObjectID(), ItemId: itemId, Name: name, UpdatedAt: time, CreatedAt: time}
	result, err := ir.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).InsertOne(ctx, &doc)

	if err != nil {
		return "", applicationerror.NewApplicationError("Fail create item.", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return "", applicationerror.NewApplicationError(fmt.Sprintf("Fail cast to ObjectId. result: %v\n", result), nil)
}

func (ir *itemRepository) FindById(ctx context.Context, id string) (*model.Item, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, applicationerror.NewApplicationError(fmt.Sprintf("argument id is not ObjectId. id: %s\n", id), err)
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	var item Item
	if err := ir.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, filter).Decode(&item); err != nil {
		return nil, applicationerror.NewApplicationError(fmt.Sprintf("Fail findById id: %s\n", id), err)
	}

	return toModel(item), nil
}

func toModel(doc Item) *model.Item {
	return &model.Item{
		Id:        doc.Id.Hex(),
		ItemId:    doc.ItemId,
		Name:      doc.Name,
		UpdatedAt: doc.UpdatedAt,
		CreatedAt: doc.CreatedAt,
	}
}

func (ir *itemRepository) FindInItemId(ctx context.Context, itemIds []string) []model.Item {
	filter := bson.D{{Key: "itmId", Value: bson.D{{Key: "$in", Value: itemIds}}}}
	cursol, err := ir.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var results []Item
	if err := cursol.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	var items []model.Item
	for _, i := range results {
		items = append(items, *toModel(i))
	}

	return items
}
