package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ayahiro1729/onpu/api/config"
)

type AuthHandler struct {
	spotifyConfig *config.SpotifyConfig
	authService  *service.AuthService
}

func NewAuthHandler(spotifyConfig *config.SpotifyConfig, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		spotifyConfig: spotifyConfig,
		authService:  authService,
	}
}

// Spotifyの認証画面にリダイレクト
func (h *AuthHandler) RedirectToSpotifyAuth(c *gin.Context) {
	authURL := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=user-read-private user-read-email",
		h.spotifyConfig.ClinetID,
		h.spotifyConfig.RedirectURI,
	)
	c,Redirect(http.StatusFound, authURL)
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
