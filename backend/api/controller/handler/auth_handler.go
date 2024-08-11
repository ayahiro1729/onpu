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

	// フロントエンドにリダイレクト
	redirectURL := fmt.Sprintf("http://localhost:3000/user/%d", userID)
	c.Redirect(http.StatusFound, redirectURL)
}

// セッションからアクセストークンを取得
func (h *AuthHandler) GetAccessTokenFromSession(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("access_token")
	fmt.Printf("Session data: %+v\n", session.Flashes()...)
  fmt.Printf("Retrieved token from session: %v\n", token)
	if token == nil {
		fmt.Println("Token not found in session")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
