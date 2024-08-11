package service

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
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
	if err := s.followPersistence.FollowUser(followerID, followeeID); err != nil {
		fmt.Printf("error following user (service): %v\n", err)
		return err
	}
	
	return nil
}
