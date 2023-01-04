package container_testcontainers

import (
	"context"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type mongoContainer struct {
	testcontainers.Container
}

func SetupMongo(ctx context.Context) (*mongoContainer, error) {
	port, _ := nat.NewPort("", "27017")
	timeout := 2 * time.Minute // default

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		Env: map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": "user",
			"MONGO_INITDB_ROOT_PASSWORD": "password",
		},
		WaitingFor: wait.ForListeningPort(port).WithStartupTimeout(timeout),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &mongoContainer{Container: container}, nil
}
