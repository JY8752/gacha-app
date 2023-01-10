package datastore

import (
	repository "JY8752/gacha-app/domain/repository/useritem"
	datastore "JY8752/gacha-app/infrastructure/datastore/mongo"
	testcontainers "JY8752/gacha-app/test/container/testcontainers"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	rep       repository.UserItemRepository
	container *testcontainers.MongoContainer
)

func startContainer() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	var err error
	container, err = testcontainers.SetupMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := datastore.NewMongoClient(&container.Client)
	rep = NewUserItemRepository(client)
}

func closeContainer() {
	container.Close(context.Background())
}

func TestUserItem(t *testing.T) {
	g := Goblin(t)

	// setup
	startContainer()
	defer closeContainer()

	g.Describe("Create", func() {
		g.It("should crate one user item", func() {
			time := time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC)
			oid, err := rep.Create(context.Background(), primitive.NewObjectID(), "item-1", 1, time)

			g.Assert(err).IsNil()
			g.Assert(oid).IsNotNil()
		})
	})

	g.Describe("FindByUserIdAndItemId", func() {
		g.It("should find one user item", func() {
			time := time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC)
			uid := primitive.NewObjectID()
			_, err := rep.Create(context.Background(), uid, "item-1", 1, time)
			if err != nil {
				g.Errorf("Fail to create document. err: %s\n", err.Error())
			}

			result, err := rep.FindByUserIdAndItemId(context.Background(), uid, "item-1")
			if err != nil {
				g.Errorf("Not found document. err: %s\n", err.Error())
			}

			g.Assert(result.UserId).Eql(uid.Hex())
			g.Assert(result.ItemId).Eql("item-1")
		})
	})

	g.Describe("IncrementCount", func() {
		oid1 := primitive.NewObjectID()
		time := time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC)

		testcases := []struct {
			name          string
			userItems     []UserItem
			userId        primitive.ObjectID
			itemId        string
			expectedCount int
			isErr         bool
		}{
			{
				name:      "when do not exist document, return error",
				userItems: []UserItem{},
				isErr:     true,
			},
			{
				name: "should increment count",
				userItems: []UserItem{
					{Id: primitive.NewObjectID(), UserId: oid1, ItemId: "item-1", Count: 1},
				},
				userId:        oid1,
				itemId:        "item-1",
				expectedCount: 2,
				isErr:         false,
			},
		}

		g.BeforeEach(func() {
			fmt.Println("------------------ BeforeEach --------------------")
			rep.Delete(context.Background())
		})

		for _, testcase := range testcases {
			testcase := testcase
			g.It(testcase.name, func() {
				// given
				for _, userItem := range testcase.userItems {
					rep.Create(context.Background(), userItem.UserId, userItem.ItemId, userItem.Count, time)
				}

				// when
				err := rep.IncrementCount(context.Background(), testcase.userId, testcase.itemId, time)

				// then
				if testcase.isErr {
					g.Assert(err.Error()).Eql("msg:  err: mongo: no documents in result\n")
				} else {
					find, err := rep.FindByUserIdAndItemId(context.Background(), testcase.userId, testcase.itemId)
					if err != nil {
						g.Errorf("Not found document. err: %s\n", err.Error())
					}
					g.Assert(find.Count).Eql(testcase.expectedCount)
				}
			})
		}
	})

	g.Describe("List", func() {
		time := time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC)
		uid := primitive.NewObjectID()

		testcases := []struct {
			name       string
			userItems  []UserItem
			expectSize int
		}{
			{
				name:       "when do not exist document, return empty slice",
				userItems:  []UserItem{},
				expectSize: 0,
			},
			{
				name: "should get one document",
				userItems: []UserItem{
					{
						UserId: uid,
						ItemId: "item-1",
					},
				},
				expectSize: 1,
			},
			{
				name: "should get two documents",
				userItems: []UserItem{
					{
						UserId: uid,
						ItemId: "item-1",
					},
					{
						UserId: uid,
						ItemId: "item-2",
					},
				},
				expectSize: 2,
			},
		}

		g.BeforeEach(func() {
			rep.Delete(context.Background())
		})

		for _, testcase := range testcases {
			testcase := testcase
			g.It(testcase.name, func() {
				for _, userItem := range testcase.userItems {
					rep.Create(context.Background(), userItem.UserId, userItem.ItemId, 1, time)
				}
				result := rep.List(context.Background(), uid)
				g.Assert(len(result)).Eql(testcase.expectSize)
			})
		}
	})
}
