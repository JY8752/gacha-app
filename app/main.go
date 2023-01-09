package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello")

	// 環境変数読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Not exist .env file.")
	}
}
