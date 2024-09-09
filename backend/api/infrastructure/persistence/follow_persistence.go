package persistence

import (
	"fmt"
	"errors"

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
		Select("users.id AS user_id, users.user_name, users.display_name, users.icon_image, follows.updated_at").
		Joins("left join users on follows.follower_id = users.id").
		Where("follows.followee_id = ? AND follows.deleted_at IS NULL", userID).
		Scan(&followers).Error; err != nil {
		fmt.Printf("error during select from follows when getting followers (persistence): %v\n", err)
		return nil, err
	}

	return &followers, nil
}

func (p *FollowPersistence) GetFollowees(userID int) (*[]repository.FollowUserDTO, error) {
	followees := []repository.FollowUserDTO{}

	if err := p.db.Model(&model.Follow{}).
		Select("users.id AS user_id, users.user_name, users.display_name, users.icon_image, follows.updated_at").
		Joins("left join users on follows.followee_id = users.id").
		Where("follows.follower_id = ? AND follows.deleted_at IS NULL", userID).
		Scan(&followees).Error; err != nil {
		fmt.Printf("error during select from follows when getting followees (persistence): %v\n", err)
		return nil, err
	}

	return &followees, nil
}

func (p *FollowPersistence) FollowUser(followerID int, followeeID int) error {
	// 既にフォロー関係がないか確認
	var existingFollow model.Follow
	err := p.db.Unscoped().Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		First(&existingFollow).Error

	if err == nil {
		if existingFollow.DeletedAt != nil {
			// 既に論理削除されたフォロー関係がある場合、deleted_atをnullにして復元
			err = p.db.Unscoped().Model(&existingFollow).Update("deleted_at", nil).Error
			if err != nil {
				fmt.Printf("error during restoring follow relationship: %v\n", err)
				return err
			}
			// 復元に成功したら、レコード追加操作は行わない
			return nil
		}
		// 既にフォローしている場合
		return fmt.Errorf("already following this user")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// レコードが存在しない(期待されている結果)以外のエラーが発生した場合
		fmt.Printf("error during checking follow relationship: %v\n", err)
		return err
	}

	follow := model.Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
	}

	// 新規フォロー関係をテーブルに追加
	if err := p.db.Select("FollowerID", "FolloweeID").Create(&follow).Error; err != nil {
		fmt.Printf("errror during creating new record to follows table: %v\n", err)
		return err
	}

	return nil
}

func (p *FollowPersistence) UnfollowUser(followerID int, followeeID int) error {
	result := p.db.Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		Delete(&model.Follow{})
	err := result.Error

	// 削除中にエラーが発生した場合
	if err != nil {
		fmt.Printf("errror during creating new record to follows table: %v\n", err)
		return err
	}

	// 削除したいレコードがなかった場合
	if result.RowsAffected == 0 {
		fmt.Printf("no follow relationship found", followerID, followeeID)
		return fmt.Errorf("no follow relationship found")
	}

	return nil
}
