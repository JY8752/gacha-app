package datastore

import (
	applicationerror "JY8752/gacha-app/error"
	r "JY8752/gacha-app/infrastructure/datastore/redis"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const KEY = "gacha_history_%s"

func getKey(gachaId string) string {
	return fmt.Sprintf(KEY, gachaId)
}

type gachaHistoryRepository struct {
	client *r.RedisClient
}

func NewGachaHistoryRepository(c *r.RedisClient) *gachaHistoryRepository {
	return &gachaHistoryRepository{c}
}

func (g *gachaHistoryRepository) Add(ctx context.Context, gachaId, itemId string, userId primitive.ObjectID) error {
	if err := g.client.HSet(ctx, getKey(gachaId), map[string]string{userId.Hex(): itemId}).Err(); err != nil {
		return applicationerror.NewApplicationError("", err)
	}
	return nil
}

func (g *gachaHistoryRepository) Get(ctx context.Context, gachaId string, userId primitive.ObjectID) string {
	result, err := g.client.HGet(ctx, getKey(gachaId), userId.Hex()).Result()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return result
}
