package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayahiro1729/onpu/api/usecase/service"

	"github.com/gin-gonic/gin"
)

type MusicListHandler struct {
	musicListService *service.MusicListService
}

func NewMusicListHandler(musicListService *service.MusicListService) *MusicListHandler {
	return &MusicListHandler{
		musicListService: musicListService,
	}
}

func (h *MusicListHandler) LatestMusicList(c *gin.Context) {
	userIDStr := c.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Printf("error reading param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	musicList, err := h.musicListService.LatestMusicList(userID)
	if err != nil {
		fmt.Printf("error getting latest music list (handler): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"musicList": musicList,
	})
}

type MusicListRequest struct {
	UserID int `json:"user_id"`
}

// ユーザーのmusic_listを作成
func (h *MusicListHandler) PostMusicList(c *gin.Context) {
	var req MusicListRequest
	//request body からuser_idを取得
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := req.UserID

	// access_tokenを取得
	token, err := h.musicListService.CheckAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// users.idを取得
	// userID, err := h.musicListService.FetchUserID(token)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// ユーザーのお気に入りの曲を10曲取得
	musics, err := h.musicListService.FetchTenFavoriteMusics(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// music_listのデータを作成
	err = h.musicListService.CreateMusicList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザーの最新のmusic_listのidを取得
	musicListID, err := h.musicListService.GetLatestMusicListID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// musicのデータを作成
	err = h.musicListService.CreateSingleMusics(musicListID, musics)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "music list created",
	})
}
