package handler

import (
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

	session := sessions.Default(c)
	session.Set("access_token", token)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
