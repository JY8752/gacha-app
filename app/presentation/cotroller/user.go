package controller

import (
	service "JY8752/gacha-app/domain/service/user"
	grpc "JY8752/gacha-app/pkg/grpc/user"
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type userController struct {
	grpc.UnimplementedUserServer
	userService service.UserService
}

func NewUserController(userService service.UserService) grpc.UserServer {
	return &userController{userService: userService}
}

func (u *userController) Create(ctx context.Context, req *grpc.CreateRequest) (*grpc.CreateResponse, error) {
	time := time.Now()
	user, err := u.userService.Create(ctx, req.Name, time)
	if err != nil {
		return nil, err
	}

	return &grpc.CreateResponse{Id: user.Id, Name: user.Name, CreatedAt: timestamppb.New(user.CreatedAt)}, nil
}

func (u *userController) ListUserItems(ctx context.Context, req *grpc.ListUserItemsRequest) (*grpc.ListUserItemsResponse, error) {
	// time := time.Now()
	return &grpc.ListUserItemsResponse{}, nil
}
