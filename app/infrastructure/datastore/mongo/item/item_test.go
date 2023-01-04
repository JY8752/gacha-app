package datastore

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	repository "JY8752/gacha-app/domain/repository/item"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	container_testcontainers "JY8752/gacha-app/test/container/testcontainers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rep repository.ItemRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, err := container_testcontainers.SetupMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	host, _ := container.Host(ctx)
	p, _ := container.MappedPort(ctx, "27017/tcp")

	connectionString := fmt.Sprintf("mongodb://user:password@%s:%d/?connect=direct", host, uint(p.Int()))
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(
		connectionString,
	))
	if err != nil {
		log.Fatal(err)
	}

	client := datastore.NewMongoClient(mongoClient)
	rep = NewItemRepository(client)

	code := m.Run()

	if err = mongoClient.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestItem(t *testing.T) {
	ctx := context.Background()
	time := time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC)

	oid, err := rep.Create(ctx, "1", "item1", time)
	if err != nil {
		t.Fatalf(err.Error())
	}

	item, err := rep.FindById(ctx, oid)
	if err != nil {
		t.Fatalf("fail find item. err: %s\n", err.Error())
	}

	if item.Id != oid {
		t.Fatalf("expect id is %s, but %s\n", oid, item.Id)
	}

	if item.ItemId != "1" {
		t.Fatalf("expect itemId is '1', but %s\n", item.ItemId)
	}

	if item.Name != "item1" {
		t.Fatalf("expect name is 'item1', but %s\n", item.Name)
	}

	if item.UpdatedAt != time {
		t.Fatalf("expect updatedAt is %v, but %v\n", time, item.UpdatedAt)
	}

	if item.CreatedAt != time {
		t.Fatalf("expect createdAt is %v, but %v\n", time, item.CreatedAt)
	}
}
