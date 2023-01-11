package controller

import (
	item_service "JY8752/gacha-app/domain/service/item"
	user_service "JY8752/gacha-app/domain/service/user"
	"JY8752/gacha-app/pkg/grpc/item"
	grpc "JY8752/gacha-app/pkg/grpc/user"
	"JY8752/gacha-app/registory"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userController struct {
	grpc.UnimplementedUserServer
	userService user_service.UserService
	itemService item_service.ItemService
}

func NewUserController(r registory.ServiceRegistory) grpc.UserServer {
	return &userController{userService: user_service.NewUserService(r), itemService: item_service.NewItemService(r)}
}

func (u *userController) Create(ctx context.Context, req *grpc.CreateRequest) (*grpc.CreateResponse, error) {
	time := time.Now()
	user, err := u.userService.Create(ctx, req.Name, time)
	if err != nil {
		return &grpc.CreateResponse{}, err
	}

	return &grpc.CreateResponse{Id: user.Id, Name: user.Name, CreatedAt: timestamppb.New(user.CreatedAt)}, nil
}

func (u *userController) ListUserItems(ctx context.Context, req *grpc.ListUserItemsRequest) (*grpc.ListUserItemsResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return &grpc.ListUserItemsResponse{}, err
	}

	result := u.userService.ListUserItems(ctx, oid)

	var itemIds []string
	var items []*item.Item

	for _, item := range result {
		itemIds = append(itemIds, item.ItemId)
	}

	itemMap := u.itemService.FindInItemId(ctx, itemIds)

	for _, ui := range result {
		if i, ok := itemMap[ui.ItemId]; ok {
			items = append(items, &item.Item{ItemId: ui.ItemId, Name: i.Name, Count: int32(ui.Count)})
		}
	}

	return &grpc.ListUserItemsResponse{Items: items}, nil
}
