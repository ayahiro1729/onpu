package service

import (
	"fmt"
	"errors"

	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"

	"gorm.io/gorm"
)

type FollowService struct {
	followPersistence persistence.FollowPersistence
}

func NewFollowService(followPersistence persistence.FollowPersistence) *FollowService {
	return &FollowService{followPersistence: followPersistence}
}

func (s *FollowService) GetFollowers(userID int) (*[]repository.FollowUserDTO, error) {
	followers, err := s.followPersistence.GetFollowers(userID)
	if err != nil {
		fmt.Printf("error getting followers (service): %v\n", err)
		return nil, err
	}

	return followers, nil
}

func (s *FollowService) GetFollowees(userID int) (*[]repository.FollowUserDTO, error) {
	followees, err := s.followPersistence.GetFollowees(userID)
	if err != nil {
		fmt.Printf("error getting followees (service): %v\n", err)
		return nil, err
	}

	return followees, nil
}

func (s *FollowService) FollowUser(followerID int, followeeID int) error {
	// TODO: APIを叩いた人が本人かSessionで確認する
	if err := s.followPersistence.FollowUser(followerID, followeeID); err != nil {
		fmt.Printf("error following user (service): %v\n", err)
		return err
	}

	return nil
}

func (s *FollowService) UnfollowUser(followerID int, followeeID int) error {
	// TODO: APIを叩いた人が本人かSessionで確認する
	if err := s.followPersistence.UnfollowUser(followerID, followeeID); err != nil {
		fmt.Printf("error unfollowing user (service): %v\n", err)
		return err
	}

	return nil
}

func (s *FollowService) GetIsFollowing(followerID int, followeeID int) (bool, error) {
	_, err := s.followPersistence.FindFollow(followerID, followeeID)
	// レコードが存在しない場合
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}