package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayahiro1729/onpu/api/usecase/service"
	
	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	followService *service.FollowService
}

func NewFollowHandler(followService *service.FollowService) *FollowHandler {
	return &FollowHandler{
		followService: followService,
	}
}

func (h *FollowHandler) GetFollowers(c *gin.Context) {
	userIDStr := c.Param("user_id")
	fmt.Printf(userIDStr)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Printf("error reading param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followers, err := h.followService.GetFollowers(userID)
	if err != nil {
		fmt.Printf("error getting followers (handler): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"followers": followers,
	})
}

func (h *FollowHandler) GetFollowees(c *gin.Context) {
	userIDStr := c.Param("user_id")
	fmt.Printf(userIDStr)

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Printf("error reading param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followees, err := h.followService.GetFollowees(userID)
	if err != nil {
		fmt.Printf("error getting followees (handler): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"followees": followees,
	})
}
