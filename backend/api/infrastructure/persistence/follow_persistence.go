package persistence

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"

	"gorm.io/gorm"
)

type FollowPersistence struct {
	db *gorm.DB
}

func NewFollowPersistence(db *gorm.DB) *FollowPersistence {
	return &FollowPersistence{db: db}
}

func (p *FollowPersistence) GetFollowers(userID int) (*[]repository.FollowUserDTO, error) {
	followers := []repository.FollowUserDTO{}

	if err := p.db.Model(&model.Follow{}).
			Select("users.id", "users.user_name", "users.user_name", "users.display_name", "users.icon_image", "follows.updated_at").
			Joins("left join users on follows.followee_id = users.id").
			Where("follows.followee_id = ?", userID).
			Scan(&followers).Error;
	err != nil {
		fmt.Printf("error during select from follows when getting followers (persistence): %v\n", err)
		return nil, err
	}

	return &followers, nil
}

func (p *FollowPersistence) GetFollowees(userID int) (*[]repository.FollowUserDTO, error) {
	followees := []repository.FollowUserDTO{}

	if err := p.db.Model(&model.Follow{}).
			Select("users.id, users.user_name, users.user_name, users.display_name, users.icon_image, follows.updated_at").
			Joins("left join users on follows.follower_id = users.id").
			Where("follows.follower_id = ?", userID).
			Scan(&followees).Error;
	err != nil {
		fmt.Printf("error during select from follows when getting followees (persistence): %v\n", err)
		return nil, err
	}

	return &followees, nil
}
