package container_testcontainers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoContainer struct {
	testcontainers.Container
	mongo.Client
}

const (
	ROOT_USERNAME   = "user"
	ROOT_PASSWORD   = "password"
	STARTUP_TIMEOUT = 2 * time.Minute // default
)

func SetupMongo(ctx context.Context) (*MongoContainer, error) {
	port, _ := nat.NewPort("", "27017")

	req := testcontainers.ContainerRequest{
		Image:        "mongo:5.0.12",
		ExposedPorts: []string{"27017/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": ROOT_USERNAME,
			"MONGO_INITDB_ROOT_PASSWORD": ROOT_PASSWORD,
		},
		WaitingFor: wait.ForListeningPort(port).WithStartupTimeout(STARTUP_TIMEOUT),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	host, _ := container.Host(ctx)
	p, _ := container.MappedPort(ctx, "27017/tcp")

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/?connect=direct", ROOT_USERNAME, ROOT_PASSWORD, host, uint(p.Int()))
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(
		connectionString,
	))

	if err != nil {
		log.Fatal(err)
	}

	return &MongoContainer{container, *mongoClient}, nil
}

func (m *MongoContainer) Close(ctx context.Context) {
	if err := m.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
