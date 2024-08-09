package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/controller/handler"
	"github.com/ayahiro1729/onpu/api/controller/middleware"
	"github.com/ayahiro1729/onpu/api/usecase/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

const (
	apiVersion = "/api/v1"
)

func NewServer() (*gin.Engine, error) {
	r := gin.Default()
	opts := middleware.ServerLogJsonOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
		Indent: 4,
	}
	loghandler := middleware.NewServerLogJsonHandler(os.Stdout, opts)
	logger := slog.New(loghandler)

	spotifyConfig, err := config.NewSpotifyConfig()
	if err != nil {
		return nil, err
	}

	// setting a CORS
	// setting a logger
	r.Use(middleware.Cors(), middleware.Logger(logger))

	tag := r.Group(apiVersion)
	{
		systemHandler := handler.NewSystemHandler()
		tag.GET("/system/health", systemHandler.Health)
	}

	log.Printf("Starting auth routing...")
	authService := service.NewAuthService(spotifyConfig)
	authHandler := handler.NewAuthHandler(authService)
	// Spotifyの認証画面にリダイレクト
	tag.GET("/spotify", authHandler.RedirectToSpotifyAuth)
	// Spotifyからのリダイレクトを受け取り、アクセストークンを取得
	tag.GET("/spotify/callback", authHandler.GetAccessTokenFromSpotify)
	fmt.Println("Auth routes have been set up.")

	log.Printf("Starting music list routing...")
	musicListService := service.NewMusicListUsecase(service.NewMusicListRepository())
	musicListHandler := handler.NewMusicListHandler(musicListService)
	tag.GET("/music/:user_id", musicListHandler.GetLatestMusicList)
	fmt.Println("Music list routes have been set up.")

	for _, route := range r.Routes() {
		fmt.Printf("Method: %s - Path: %s\n", route.Method, route.Path)
	}

	return r, nil
}
