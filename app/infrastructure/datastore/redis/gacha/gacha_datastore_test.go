package datastore_test

import (
	repository "JY8752/gacha-app/domain/repository/gacha"
	"JY8752/gacha-app/infrastructure/datastore/redis"
	datastore "JY8752/gacha-app/infrastructure/datastore/redis/gacha"
	container_testcontainers "JY8752/gacha-app/test/container/testcontainers"
	"context"
	"log"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var rep repository.GachaHistoryRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, err := container_testcontainers.SetupRedis(ctx)
	if err != nil {
		log.Fatal(err)
	}

	c := redis.NewRedisClient(container.Client)
	rep = datastore.NewGachaHistoryRepository(c)

	code := m.Run()

	container.FlushAll(ctx)
	os.Exit(code)
}

func TestGachaHistoryRepository(t *testing.T) {
	ctx := context.Background()
	oid := primitive.NewObjectID()

	if err := rep.Add(ctx, "gacha-1", "item-1", oid); err != nil {
		t.Fatal(err)
	}

	result := rep.Get(ctx, "gacha-1", oid)

	if result != "item-1" {
		t.Fatalf("expected 'item-1', but result: %s\n", result)
	}
}

func TestGetNoValue(t *testing.T) {
	ctx := context.Background()
	oid := primitive.NewObjectID()

	result := rep.Get(ctx, "gacha-1", oid)

	if result != "" {
		t.Fatalf("expected '', but result: %s\n", result)
	}
}
