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
	UserID interface{} `json:"user_id"`
}

// ユーザーのmusic_listを作成
func (h *MusicListHandler) PostMusicList(c *gin.Context) {
	var req MusicListRequest
	//request body からuser_idを取得
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userID int
    switch v := req.UserID.(type) {
    case float64:
        userID = int(v)
    case string:
        var err error
        userID, err = strconv.Atoi(v)
        if err != nil {
            fmt.Printf("Error converting user_id to int: %v\n", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
            return
        }
    default:
        fmt.Printf("Unexpected type for user_id: %T\n", v)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
        return
    }

	// access_tokenを取得
	token, err := h.musicListService.CheckAccessToken()
	if err != nil {
		fmt.Printf("Error checking access token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Access token retrieved successfully\n")

	// ユーザーのお気に入りの曲を10曲取得
	musics, err := h.musicListService.FetchTenFavoriteMusics(token)
	if err != nil {
		fmt.Printf("Error fetching favorite musics: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Retrieved %d favorite musics\n", len(musics))

	// music_listのデータを作成
	err = h.musicListService.CreateMusicList(userID)
	if err != nil {
		fmt.Printf("Error creating music list: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Music list created for user ID: %d\n", userID)

	// ユーザーの最新のmusic_listのidを取得
	musicListID, err := h.musicListService.GetLatestMusicListID(userID)
	if err != nil {
		fmt.Printf("Error getting latest music list ID: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Latest music list ID: %d\n", musicListID)

	// musicのデータを作成
	err = h.musicListService.CreateSingleMusics(musicListID, musics)
	if err != nil {
		fmt.Printf("Error creating single musics: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Single musics created successfully\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "music list created",
	})
}
