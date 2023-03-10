package service

import (
	model "JY8752/gacha-app/domain/model/useritem"
	"JY8752/gacha-app/registory"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GachaService interface {
	Buy(ctx context.Context, userId primitive.ObjectID, gachaId string, time time.Time) (*model.UserItem, error)
}

type gachaService struct {
	registory.ServiceRegistory
}

func NewGachaService(r registory.ServiceRegistory) *gachaService {
	return &gachaService{r}
}

func (g *gachaService) Buy(ctx context.Context, userId primitive.ObjectID, gachaId string, time time.Time) (*model.UserItem, error) {
	// 指定のガチャ筐体を取得する
	gacha, err := g.Gacha().FindByGachaId(ctx, gachaId)
	if err != nil {
		return nil, err
	}

	// 前回取得アイテムと被っていたらやりなおし
	history := g.GachaHistory().Get(ctx, gachaId, userId)

	var itemId string
	for {
		// ガチャ筐体のアイテム一覧から抽選する
		if itemId = gacha.Lottery(); itemId != history {
			break
		}
	}

	// ユーザーアイテムを更新する
	if err := g.UserItem().IncrementCount(ctx, userId, itemId, time); err != nil {
		// レコードがまだ存在していないので新規でインサートする
		_, err = g.UserItem().Create(ctx, userId, itemId, 1, time)
		if err != nil {
			return nil, err
		}
	}

	// 取得アイテムを記録
	if err = g.GachaHistory().Add(ctx, gachaId, itemId, userId); err != nil {
		return nil, err
	}

	// レスポンス返す
	item, err := g.UserItem().FindByUserIdAndItemId(ctx, userId, itemId)
	if err != nil {
		return nil, err
	}
	return item, nil
}
