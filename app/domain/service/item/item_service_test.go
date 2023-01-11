package service_test

import (
	model "JY8752/gacha-app/domain/model/item"
	repository "JY8752/gacha-app/domain/repository/item"
	service "JY8752/gacha-app/domain/service/item"
	"JY8752/gacha-app/registory"
	"context"
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

// mock

type mockItemServiceRegistory struct {
	registory.ServiceRegistory
	mockItemRepository
}

func (isr *mockItemServiceRegistory) Item() repository.ItemRepository {
	return &isr.mockItemRepository
}

type mockItemRepository struct {
	repository.ItemRepository
	fakeFindInItemId func(context.Context, []string) []model.Item
}

func (ir *mockItemRepository) FindInItemId(ctx context.Context, ids []string) []model.Item {
	return ir.fakeFindInItemId(ctx, ids)
}

func TestFindInItemId(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("FindInItemId", func() {
		testcases := []struct {
			name             string
			fakeFindInItemId func(context.Context, []string) []model.Item
			expect           map[string]model.Item
		}{
			{
				"Find item success",
				func(ctx context.Context, ids []string) []model.Item {
					return []model.Item{
						{ItemId: "item-1", Name: "item-A"}, {ItemId: "item-2", Name: "item-B"},
					}
				},
				map[string]model.Item{
					"item-1": {ItemId: "item-1", Name: "item-A"},
					"item-2": {ItemId: "item-2", Name: "item-B"},
				},
			},
			{
				"Find item success: no item",
				func(ctx context.Context, ids []string) []model.Item {
					return []model.Item{}
				},
				map[string]model.Item{},
			},
		}

		for _, testcase := range testcases {
			testcase := testcase

			g.It(testcase.name, func() {
				r := &mockItemServiceRegistory{}
				r.fakeFindInItemId = testcase.fakeFindInItemId
				s := service.NewItemService(r)

				m := s.FindInItemId(context.Background(), []string{})
				fmt.Println(m)
				g.Assert(m).Eql(testcase.expect)
			})
		}
	})
}
