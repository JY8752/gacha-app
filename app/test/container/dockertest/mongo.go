package container_dockertest

import (
	"context"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoContainer struct {
	DbClient *mongo.Client

	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func Start() (*MongoContainer, error) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Printf("Could not construct pool: %s\n", err)
		return nil, err
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Printf("Could not connect to Docker: %s", err)
		return nil, err
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=user",
			"MONGO_INITDB_ROOT_PASSWORD=password",
		},
	}

	resource, err := pool.RunWithOptions(runOptions,
		func(hc *docker.HostConfig) {
			hc.AutoRemove = true
			hc.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)

	port := resource.GetPort("27017/tcp")

	var dbClient *mongo.Client
	pool.Retry(func() error {
		dbClient, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://user:password@localhost:%s", port),
			),
		)
		if err != nil {
			return err
		}
		return dbClient.Ping(context.TODO(), nil)
	})

	if err != nil {
		log.Printf("Could not connect to docker: %s", err)
		return nil, err
	}

	fmt.Println("start mongo container🐳")

	return &MongoContainer{DbClient: dbClient, pool: pool, resource: resource}, nil
}

func (m *MongoContainer) Close() (err error) {
	// When you're done, kill and remove the container
	if err = m.pool.Purge(m.resource); err != nil {
		log.Printf("Could not purge resource: %s", err)
		return err
	}

	// disconnect mongodb client
	if err = m.DbClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("close mongo container🐳")

	return nil
}
