package main

import (
	"context"
	"fmt"

	"github.com/ayahiro1729/onpu/api/controller"
	"github.com/ayahiro1729/onpu/api/infrastructure/database"
)

func main() {
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run()

	ctx := context.Background()
	db := database.New(ctx)
	if db != nil {
		fmt.Println("PostgreSQLに接続成功")
	} else {
		fmt.Println("PostgreSQLに接続失敗")
	}
}
