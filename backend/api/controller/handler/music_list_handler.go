package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayahiro1729/onpu/api/usecase/service"

	"github.com/gin-contrib/sessions"
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

// ユーザーのmusic_listを作成
func (h *MusicListHandler) PostMusicList(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// セッションからaccess_tokenを取得
	sessionAccessToken := sessions.DefaultMany(c, "access_token")
	accessToken, ok := sessionAccessToken.Get("access_token").(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not found"})
		return
	}

	// Spotifyから最近よく聴く10曲を取得
	tracks, err := h.musicListService.GetTopTracks(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get top tracks from Spotify"})
		return
	}

	// 新しいMusicListをDBに保存
	musicList, err := h.musicListService.CreateMusicList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create music list"})
		return
	}

	// tracksを一つずつ新しいmusicとしてDBに保存
	if err := h.musicListService.CreateMusics(musicList.ID, tracks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create musics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Music list created successfully",
		"musicList": musicList,
	})
}
