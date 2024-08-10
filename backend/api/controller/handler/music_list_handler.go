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
	fmt.Printf(userIDStr)

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
