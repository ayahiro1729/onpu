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

	session := sessions.Default(c)
	session.Set("access_token", token)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザー情報を取得
	user, err := h.authService.FetchSpotifyUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// DBからuser_idを取得（アカウントがない場合は登録）
	user_id, err := h.authService.FetchUserID(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// フロントエンドにリダイレクト
	redirectURL := fmt.Sprintf("http://localhost:3000/user/%d", user_id)
	c.Redirect(http.StatusFound, redirectURL)
}
