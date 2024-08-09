package controller

import (
	"fmt"
	"os"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/controller/handler"
	"github.com/ayahiro1729/onpu/api/controller/middleware"
	"github.com/ayahiro1729/onpu/api/usecase/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// setting a session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	tag := r.Group(apiVersion)
	// ヘルスチェックAPI
	{
		systemHandler := handler.NewSystemHandler()

		tag.GET("/system/health", systemHandler.Health)
	}

	// Spotify認証API
	{
		authService := service.NewAuthService(spotifyConfig)
		authHandler := handler.NewAuthHandler(authService)

		// Spotifyからのリダイレクトを受け取り、アクセストークンを取得
		tag.POST("/user", authHandler.GetAccessTokenFromSpotify)
	}

	for _, route := range r.Routes() {
		fmt.Printf("Method: %s - Path: %s\n", route.Method, route.Path)
	}

	return r, nil
}
