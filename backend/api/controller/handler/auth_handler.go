package handler

import (
	"net/http"

	"github.com/ayahiro1729/onpu/api/usecase/service"
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

// Spotifyの認証画面にリダイレクト
func (h *AuthHandler) RedirectToSpotifyAuth(c *gin.Context) {
	authURL := h.authService.GetSpotifyAuthURL()
	c.Redirect(http.StatusFound, authURL)
}

// Spotifyからのリダイレクトを受け取り、アクセストークンを取得
func (h *AuthHandler) GetAccessTokenFromSpotify(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided"})
		return
	}

	// 認可コードを使用してアクセストークンを取得
	token, err := h.authService.GetSpotifyToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// トークンを使用してSpotifyのユーザ情報を取得
	user, err := h.authService.GetSpotifyUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンを生成
	jwtToken, err := h.authService.GenerateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
