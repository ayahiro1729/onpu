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
	fmt.Println(userIDStr)

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
	fmt.Println(userIDStr)

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

func (h *FollowHandler) FollowUser(c *gin.Context) {
	followerIDStr := c.Param("follower_id")
	followeeIDStr := c.Param("followee_id")

	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		fmt.Printf("error reading follower_id param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followeeID, err := strconv.Atoi(followeeIDStr)
	if err != nil {
		fmt.Printf("error reading followee_id param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.followService.FollowUser(followerID, followeeID); err != nil {
		fmt.Printf("error following user (handler): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully followed user"})
}

func (h *FollowHandler) UnfollowUser(c *gin.Context) {
	followerIDStr := c.Param("follower_id")
	followeeIDStr := c.Param("followee_id")

	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		fmt.Printf("error reading follower_id param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followeeID, err := strconv.Atoi(followeeIDStr)
	if err != nil {
		fmt.Printf("error reading followee_id param: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.followService.UnfollowUser(followerID, followeeID); err != nil {
		fmt.Printf("error unfollowing user (handler): %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully unfollowed user"})
}

func (h *FollowHandler) GetIsFollowing(c *gin.Context) {
	followerIDStr := c.Param("follower_id")
	followeeIDStr := c.Param("followee_id")

	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followeeID, err := strconv.Atoi(followeeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isFollowing, err := h.followService.GetIsFollowing(followerID, followeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_following": isFollowing,
	})
}