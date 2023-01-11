package datastore

import (
	"JY8752/gacha-app/constant"
	model "JY8752/gacha-app/domain/model/gacha"
	item_model "JY8752/gacha-app/domain/model/item"
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

const COLLECTION_NAME = "Gachas"

type Gacha struct {
	Id        primitive.ObjectID `bson:"_id"`
	GachaId   string             `bson:"gchId"`
	Name      string             `bson:"nm"`
	Items     []Item             `bson:"itms"`
	UpdatedAt time.Time          `bson:"updAt"`
	CreatedAt time.Time          `bson:"crtAt"`
}

type Item struct {
	ItemId string `bson:"itmId"`
	Weight int    `bson:"wgh"`
}

type gachaRepository struct {
	client *datastore.MongoClient
}

func NewGachaRepository(c *datastore.MongoClient) *gachaRepository {
	return &gachaRepository{c}
}

func (g *gachaRepository) FindByGachaId(ctx context.Context, gachaId string) (*model.Gacha, error) {
	filter := bson.D{{Key: "gchId", Value: gachaId}}
	result := g.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, applicationerror.NewApplicationError("", result.Err())
		}
		log.Fatal(result.Err())
	}

	var gacha Gacha
	if err := result.Decode(&gacha); err != nil {
		log.Fatal(err)
	}

	return toModel(gacha), nil
}

func toModel(g Gacha) *model.Gacha {
	var items []item_model.ItemWith
	for _, item := range g.Items {
		items = append(items, item_model.ItemWith{ItemId: item.ItemId, Weight: item.Weight})
	}
	return &model.Gacha{
		Id:        g.Id.Hex(),
		GachaId:   g.GachaId,
		Name:      g.Name,
		Items:     items,
		UpdatedAt: g.UpdatedAt,
		CreatedAt: g.CreatedAt,
	}
}

// for test
func (g *gachaRepository) Create(ctx context.Context, gachaId, name string, items []item_model.ItemWith, time time.Time) (primitive.ObjectID, error) {
	doc := Gacha{
		Id:        primitive.NewObjectID(),
		GachaId:   gachaId,
		Name:      name,
		Items:     toEntities(items),
		UpdatedAt: time,
		CreatedAt: time,
	}

	result, err := g.client.GetDB(constant.MONGO_MAIN_DB).Collection(COLLECTION_NAME).InsertOne(ctx, &doc)
	if err != nil {
		return primitive.NilObjectID, applicationerror.NewApplicationError("", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, applicationerror.NewApplicationError(fmt.Sprintf("fail cast to ObjectId. result: %v\n", result), nil)
}

func toEntities(items []item_model.ItemWith) (entities []Item) {
	for _, i := range items {
		entities = append(entities, Item{
			ItemId: i.ItemId,
			Weight: i.Weight,
		})
	}
	return entities
}
