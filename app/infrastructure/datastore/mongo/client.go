package datastore

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(c *mongo.Client) *MongoClient {
	return &MongoClient{client: c}
}

func (mc *MongoClient) GetDB(name string) *mongo.Database {
	return mc.client.Database(name)
}

func (mc *MongoClient) Connect(ctx context.Context) error {
	return mc.client.Connect(ctx)
}

func (mc *MongoClient) Disconnect(ctx context.Context) error {
	return mc.client.Disconnect(ctx)
}
