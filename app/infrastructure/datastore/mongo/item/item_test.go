package datastore

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	repository "JY8752/gacha-app/domain/repository/item"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	container_testcontainers "JY8752/gacha-app/test/container/testcontainers"
)

var rep repository.ItemRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, err := container_testcontainers.SetupMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := datastore.NewMongoClient(&container.Client)
	rep = NewItemRepository(client)

	code := m.Run()

	container.Close(ctx)
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

func TestFindInItemId(t *testing.T) {
	ctx := context.Background()
	time := time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)
	rep.Create(ctx, "item-1", "item-1", time)
	rep.Create(ctx, "item-2", "item-2", time)

	results := rep.FindInItemId(context.Background(), []string{"item-1", "item-2"})

	if len(results) != 2 {
		t.Fatalf("expect result length is 2, but %d\n", len(results))
	}
}
