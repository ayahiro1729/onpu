package handler

import (
	"strconv"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/usecase/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// ユーザーを作成
func (h *UserHandler) PostUser(c *gin.Context) {
	// ユーザー情報を取得
	spotifyUser := &model.User{}
	if err := c.BindJSON(spotifyUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// ユーザー登録またはログイン
	user, err := h.userService.RegisterOrLogin(spotifyUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user_id": user.ID,
	})
}

// ユーザーの情報を取得（プロフィール画面）
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	id_str := c.Param("user_id")
	if id_str == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	id_int, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user_id := uint(id_int)

	// セッションからユーザー情報を取得
	user, err := h.userService.FindUserProfile(user_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

// ユーザーの情報を更新
func (h *UserHandler) PutUserProfile(c *gin.Context) {
	id_str := c.Param("user_id")
	if id_str == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	id_int, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user_id := uint(id_int)

	// リクエストボディからユーザー情報を取得
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// ユーザーIDを設定
	user.ID = int(user_id)

	// ユーザー情報を更新
	err = h.userService.UpdateUserProfile(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
