package service_test

import (
	user_model "JY8752/gacha-app/domain/model/user"
	user_repository "JY8752/gacha-app/domain/repository/user"
	useritem_repository "JY8752/gacha-app/domain/repository/useritem"
	service "JY8752/gacha-app/domain/service/user"
	"JY8752/gacha-app/registory"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/franela/goblin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// mock

// UserServiceRegistory
type mockUserServiceRegistory struct {
	registory.ServiceRegistory // インターフェイス埋め込みすることでメソッド実装なくてもコンパイル通る
	mockUserRepository
	mockUserItemRepository
}

func (musr *mockUserServiceRegistory) User() user_repository.UserRepository {
	return &musr.mockUserRepository
}

func (musr *mockUserServiceRegistory) UserItem() useritem_repository.UserItemRepository {
	return &musr.mockUserItemRepository
}

// UserRepository
type mockUserRepository struct {
	user_repository.UserRepository
	fakeCreate func(name string, time time.Time) (string, error)
}

func (mur *mockUserRepository) Create(ctx context.Context, name string, time time.Time) (string, error) {
	return mur.fakeCreate(name, time)
}

// UserItemRepository
type mockUserItemRepository struct {
	useritem_repository.UserItemRepository
}

func TestCreate(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Create", func() {
		oid := primitive.NewObjectID()
		testTime := time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC)
		testcases := []struct {
			name       string
			fakeCreate func(string, time.Time) (string, error)
			userName   string
			expect     *user_model.User
			isError    bool
		}{
			{
				name: "create one user",
				fakeCreate: func(s string, t time.Time) (string, error) {
					return oid.Hex(), nil
				},
				userName: "user1",
				expect: &user_model.User{
					Id:        oid.Hex(),
					Name:      "user1",
					UpdatedAt: testTime,
					CreatedAt: testTime,
				},
				isError: false,
			},
			{
				name: "when create error, return nil",
				fakeCreate: func(s string, t time.Time) (string, error) {
					return "", errors.New("")
				},
				userName: "user1",
				expect:   nil,
				isError:  true,
			},
		}

		for _, testcase := range testcases {
			testcase := testcase
			g.It(testcase.name, func() {
				// given
				r := &mockUserServiceRegistory{}
				r.fakeCreate = testcase.fakeCreate
				s := service.NewUserService(r)

				// when
				result, err := s.Create(context.Background(), testcase.userName, testTime)

				// then
				fmt.Printf("result: %v, err: %v\n", result, err)
				if testcase.isError {
					g.Assert(err).IsNotNil()
				} else {
					g.Assert(err).IsNil()
				}
				g.Assert(result).Eql(testcase.expect)
			})
		}
	})
}
