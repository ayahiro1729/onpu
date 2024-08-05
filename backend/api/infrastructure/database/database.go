package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ayahiro1729/onpu/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(ctx context.Context) *gorm.DB {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// PostgresのDSN (Data Source Name) を構築
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.Port,
	)

	// 環境変数に応じて接続設定を変更（必要に応じて）
	// if cfg.Env == "PROD" {
	// }

	// Postgresのダイヤレクタを作成
	dialector := postgres.Open(dsn)

	if db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		connect(dialector, 100)
	}

	return db
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		panic(err.Error())
	}
}
