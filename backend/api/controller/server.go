package controller

import (
	"fmt"
	"os"
	"net/http"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/controller/handler"
	"github.com/ayahiro1729/onpu/api/controller/middleware"
	"github.com/ayahiro1729/onpu/api/infrastructure/database"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
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

	// setting a session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
    Path:     "/",
    MaxAge:   86400 * 7, // 1週間
    HttpOnly: true,
    Secure:   false, // 開発環境ではfalse、本番環境ではtrueに
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("mysession", store))

	r.Use(middleware.SessionDebug())

	opts := middleware.ServerLogJsonOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
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

		// セッションからアクセストークンを取得
		tag.GET("/session/token", authHandler.GetAccessTokenFromSession)
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

	// DBから最新のmusic listを取得するAPI
	{
		musicListPersistence := persistence.NewMusicListPersistence(db)
		musicListService := service.NewMusicListService(*musicListPersistence)
		musicListHandler := handler.NewMusicListHandler(musicListService)

		tag.GET("/music/:user_id", musicListHandler.LatestMusicList)

		// ユーザーのmusic listを作成
		tag.POST("/music", musicListHandler.PostMusicList)
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
	}

	for _, route := range r.Routes() {
		fmt.Printf("Method: %s - Path: %s\n", route.Method, route.Path)
	}

	return r, nil
}
