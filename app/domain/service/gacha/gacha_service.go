package service

import (
	model "JY8752/gacha-app/domain/model/item"
	"JY8752/gacha-app/registory"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GachaService interface {
}

type gachaService struct {
	registory.ServiceRegistory
}

func NewGachaService(r registory.ServiceRegistory) *gachaService {
	return &gachaService{r}
}

func (g *gachaService) Buy(ctx context.Context, userId primitive.ObjectID, gachaId string) (*model.Item, error) {
	// 指定のガチャ筐体を取得する
	_, err := g.Gacha().FindByGachaId(ctx, gachaId)
	if err != nil {
		return nil, err
	}

	// ガチャ筐体のアイテム一覧から抽選する

	// 前回取得アイテムと被っていたらやりなおし

	// ユーザーアイテムを更新する

	// レスポンス返す
	return &model.Item{}, nil
}
