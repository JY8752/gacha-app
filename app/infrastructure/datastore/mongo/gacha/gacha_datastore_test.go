package datastore_test

import (
	gacha_model "JY8752/gacha-app/domain/model/gacha"
	item_model "JY8752/gacha-app/domain/model/item"
	repository "JY8752/gacha-app/domain/repository/gacha"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	gacha_datastore "JY8752/gacha-app/infrastructure/datastore/mongo/gacha"
	container_testcontainers "JY8752/gacha-app/test/container/testcontainers"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/franela/goblin"
)

var rep repository.GachaRepository

func TestMain(m *testing.M) {
	ctx := context.Background()
	container, err := container_testcontainers.SetupMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := datastore.NewMongoClient(&container.Client)
	rep = gacha_datastore.NewGachaRepository(client)

	code := m.Run()

	container.Close(ctx)
	os.Exit(code)
}

func TestFindByGachaId(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("TestFindByGachaId", func() {
		g.It("should create and find", func() {
			// given
			ctx := context.Background()
			time := time.Date(2023, 1, 11, 0, 0, 0, 0, time.UTC)
			items := []item_model.ItemWith{
				{ItemId: "item-ID-1", Weight: 10},
				{ItemId: "item-ID-2", Weight: 10},
			}

			oid, err := rep.Create(ctx, "gacha-ID-1", "gacha-1", items, time)
			if err != nil {
				g.Fail(err)
			}

			// when
			result, err := rep.FindByGachaId(ctx, "gacha-ID-1")
			if err != nil {
				g.Fail(err)
			}

			g.Assert(result).Eql(&gacha_model.Gacha{
				Id:      oid.Hex(),
				GachaId: "gacha-ID-1",
				Name:    "gacha-1",
				Items: []item_model.ItemWith{
					{ItemId: "item-ID-1", Weight: 10},
					{ItemId: "item-ID-2", Weight: 10},
				},
				UpdatedAt: time,
				CreatedAt: time,
			})
		})
	})
}
