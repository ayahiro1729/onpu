// 認証が必要なAPIのルーティング
package router

func SetAuthRouting(router *gin.Engine, authHandler *handler.AuthHandler) {
	authRoutes := router.Group("/auth")
	{
		// Spotifyの認証画面にリダイレクト
		authRoutes.GET("/auth/spotify", authHandler.RedirectToSpotifyAuth)
		// Spotifyからのリダイレクトを受け取り、アクセストークンを取得
		authRoutes.GET("/auth/spotify/callback", authHandler.GetAccessTokenFromSpotify)
	}
}
