package controller

import (
	service "JY8752/gacha-app/domain/service/gacha"
	"JY8752/gacha-app/pkg/grpc/gacha"
	"JY8752/gacha-app/pkg/grpc/item"
	"JY8752/gacha-app/registory"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type gachaController struct {
	gacha.UnimplementedGachaServer
	gachaService service.GachaService
}

func NewGachaController(r registory.ServiceRegistory) *gachaController {
	return &gachaController{gachaService: service.NewGachaService(r)}
}

func (g *gachaController) Buy(ctx context.Context, req *gacha.BuyRequest) (*gacha.BuyResponse, error) {
	time := time.Now()
	oid, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return &gacha.BuyResponse{}, err
	}

	i, err := g.gachaService.Buy(ctx, oid, req.GachaId, time)
	if err != nil {
		return &gacha.BuyResponse{}, err
	}

	return &gacha.BuyResponse{
		Item: &item.Item{
			ItemId: i.ItemId,
			Name:   "", // TODO
			Count:  int32(i.Count),
		},
	}, nil
}
