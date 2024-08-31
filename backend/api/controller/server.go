package controller

import (
	"fmt"
	"os"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/controller/handler"
	"github.com/ayahiro1729/onpu/api/controller/middleware"
	"github.com/ayahiro1729/onpu/api/infrastructure/database"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
	"github.com/ayahiro1729/onpu/api/usecase/service"

	"github.com/gin-contrib/cors"
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
	// r.Use(middleware.Cors(), middleware.Logger(logger))

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	r.Use(middleware.Logger(logger))

	// setting a session
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"access_token", "user_id"}
	// r.Use(sessions.Sessions("mysession", store))
	r.Use(sessions.SessionsMany(sessionNames, store))

	// setting a database
	db := database.NewDB()
	if db != nil {
		fmt.Println("PostgreSQLに接続成功")
	} else {
		fmt.Println("PostgreSQLに接続失敗")
	}

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

		// Spotifyからのリダイレクトを受け取り、①アクセストークンを取得、②ユーザー情報を取得、③登録またはログイン
		r.GET("/callback", authHandler.AuthenticateUser)
		tag.GET("/myuserid", authHandler.GetUserIDFromSession)
	}

	// ユーザー情報API
	{
		userRepository := repository.NewUserRepository(db)
		userService := service.NewUserService(*userRepository)
		userHandler := handler.NewUserHandler(userService)

		// ユーザーを作成
		tag.POST("/user", userHandler.PostUser)

		// ユーザーの情報を取得（プロフィール画面）
		tag.GET("/user/:user_id", userHandler.GetUserProfile)
	}

	// music list情報API
	{
		musicListPersistence := persistence.NewMusicListPersistence(db)
		musicListService := service.NewMusicListService(*musicListPersistence)
		musicListHandler := handler.NewMusicListHandler(musicListService)

		tag.GET("/music/:user_id", musicListHandler.LatestMusicList)
	}

	// フォロー情報API
	{
		followPersistence := persistence.NewFollowPersistence(db)
		followService := service.NewFollowService(*followPersistence)
		followHandler := handler.NewFollowHandler(followService)

		// あるユーザーのフォロワーを取得
		tag.GET("/follower/:user_id", followHandler.GetFollowers)

		// あるユーザーのフォロー中ユーザーを取得
		tag.GET("/followee/:user_id", followHandler.GetFollowees)

		// ユーザーをフォローする
		tag.POST("/follow/:follower_id/:followee_id", followHandler.FollowUser)

		// ユーザーのフォローを外す
		tag.DELETE("/follow/:follower_id/:followee_id", followHandler.UnfollowUser)
	}

	for _, route := range r.Routes() {
		fmt.Printf("Method: %s - Path: %s\n", route.Method, route.Path)
	}

	return r, nil
}
