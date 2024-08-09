package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ayahiro1729/onpu/api/controller"
	"github.com/ayahiro1729/onpu/api/infrastructure/database"
)

func main() {
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}

	fmt.Println("サーバーを起動します")

	if err := s.Run(":8080"); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}

	ctx := context.Background()
	db := database.New(ctx)
	if db != nil {
		fmt.Println("PostgreSQLに接続成功")
	} else {
		fmt.Println("PostgreSQLに接続失敗")
	}
}
