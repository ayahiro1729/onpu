package handler

import (
	"fmt"
	"net/http"

	"github.com/ayahiro1729/onpu/api/usecase/service"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Spotifyからのリダイレクトを受け取り、アクセストークンを取得
func (h *AuthHandler) AuthenticateUser(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
		return
	}

	// 認可コードを使用してアクセストークンを取得
	token, err := h.authService.FetchSpotifyToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sessionAccessToken := sessions.DefaultMany(c, "access_token")
	sessionAccessToken.Set("access_token", token)
	if err := sessionAccessToken.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Session saved successfully. Token: %v\n", token)

	// ユーザー情報を取得
	user, err := h.authService.FetchSpotifyUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// DBからusers.idを取得（アカウントがない場合は登録）
	userID, err := h.authService.FetchUserID(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// user_idをSessionに保存
	sessionUserID := sessions.DefaultMany(c, "user_id")
	sessionUserID.Set("user_id", userID)
	if err := sessionUserID.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(200, gin.H{
	// 	"token": sessionAccessToken.Get("access_token"),
	// 	"user_id": sessionUserID.Get("user_id"),
	// })

	// フロントエンドにリダイレクト
	redirectURL := fmt.Sprintf("http://localhost:3000/user/%d", userID)
	c.Redirect(http.StatusFound, redirectURL)
}

func (h *AuthHandler) GetUserIDFromSession(c *gin.Context) {
	sessionUserID := sessions.DefaultMany(c, "user_id")
	userID := sessionUserID.Get("user_id")

	if userID == nil {
		fmt.Println("Error getting my user id from session:")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println("Retrieved user id from session:", userID)

	c.JSON(200, gin.H{
		"user_id": sessionUserID.Get("user_id"),
	})

}

// セッションからアクセストークンを取得
func (h *AuthHandler) GetAccessTokenFromSession(c *gin.Context) {
	sessionAccessToken := sessions.DefaultMany(c, "access_token")
	token := sessionAccessToken.Get("access_token")

	if token == nil {
		fmt.Println("Error getting my token from session:")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{"access_token": sessionAccessToken.Get("access_token")})
}
