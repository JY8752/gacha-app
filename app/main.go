package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"JY8752/gacha-app/config"
	mongoclient "JY8752/gacha-app/infrastructure/datastore/mongo"
	redisclient "JY8752/gacha-app/infrastructure/datastore/redis"
	register "JY8752/gacha-app/presentation/cotroller"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Not exist .env file.")
	}

	// 8080ポートのリスナーを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// config
	config := config.GetConfig()

	// mongo
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.Uri))
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := mongoclient.NewMongoClient(client)

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient := redisclient.NewRedisClient(rdb)

	// controllerの登録
	register.RegisterController(s, mongoClient, redisClient)

	// サーバーリフレクションの設定
	reflection.Register(s)

	// サーバー起動
	go func() {
		log.Println("start gRPC server!!")
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
