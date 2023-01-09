package datastore

import (
	repository "JY8752/gacha-app/domain/repository/user"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	container "JY8752/gacha-app/test/container/dockertest"
	"context"
	"log"
	"os"
	"testing"
	"time"
)

var rep repository.UserRepository

func TestMain(m *testing.M) {
	mongoClient, close, err := container.Start()
	if err != nil {
		log.Fatal(err)
	}

	client := datastore.NewMongoClient(mongoClient)

	rep = NewUserRepository(client)

	code := m.Run()

	close()

	os.Exit(code)
}

// func TestMain(m *testing.M) {
// 	ctx := context.Background()
// 	container, err := container_testcontainers.SetupMongo(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	host, _ := container.Host(ctx)
// 	p, _ := container.MappedPort(ctx, "27017/tcp")

// 	connectionString := fmt.Sprintf("mongodb://user:password@%s:%d/?connect=direct", host, uint(p.Int()))
// 	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(
// 		connectionString,
// 	))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := datastore.NewMongoClient(mongoClient)
// 	rep = NewUserRepository(client)

// 	code := m.Run()

// 	if err = mongoClient.Disconnect(ctx); err != nil {
// 		log.Fatal(err)
// 	}

// 	os.Exit(code)
// }

func TestUser(t *testing.T) {
	time := time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC)

	id, err := rep.Create(context.Background(), "user", time)
	if err != nil {
		t.Fatalf("fail create user. err: %s\n", err.Error())
	}

	user, err := rep.FindById(context.Background(), id)
	if err != nil {
		t.Fatalf("fail find user. err: %s\n", err.Error())
	}

	if user.Id != id {
		t.Fatalf("expect id is %s, but %s\n", id, user.Id)
	}

	if user.Name != "user" {
		t.Fatalf("expect name is 'user', but %s\n", user.Name)
	}

	if user.UpdatedAt != time {
		t.Fatalf("expect updatedAt is %v, but %v\n", time, user.UpdatedAt)
	}

	if user.CreatedAt != time {
		t.Fatalf("expect createdAt is %v, but %v\n", time, user.CreatedAt)
	}
}
