package persistence

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"

	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) *UserPersistence {
	return &UserPersistence{db: db}
}

// id でユーザーを検索
func (p *UserPersistence) FindUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := p.db.First(&user, id).Error; err != nil {
		fmt.Printf("error during select from users when finding user by user_id (persistence): %v\n", err)
		return nil, err
	}
	return &user, nil
}

// spotify_idでユーザーを検索
func (p *UserPersistence) FindUserBySpotifyID(spotifyID string) (*model.User, error) {
	var user model.User
	if err := p.db.Where("spotify_id = ?", spotifyID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// user_nameでユーザーを検索
func (p *UserPersistence) FindUsersByUserName(search_string string) (*[]repository.UserSearchResultDTO, error) {
	users := []repository.UserSearchResultDTO{}

	if err := p.db.Model(&model.User{}).
		Select("id AS user_id", "user_name", "display_name", "icon_image").
		Where("user_name LIKE ?", "%"+search_string+"%").
		Scan(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// 新しいユーザーを作成
func (p *UserPersistence) CreateUser(user *model.User) error {
	if err := p.db.Create(user).Error; err != nil {
		fmt.Printf("error during create user (persistence): %v\n", err)
		return err
	}
	return nil
}

// ユーザーを更新
func (p *UserPersistence) UpdateUser(user *model.User) error {
	// ユーザーが存在するか確認
	var existingUser model.User
	if err := p.db.First(&existingUser, user.ID).Error; err != nil {
		fmt.Printf("error during select from users when updating user (persistence): %v\n", err)
		return err
	}

	// レコードを更新
	if err := p.db.Model(existingUser).Updates(user).Error; err != nil {
		fmt.Printf("error during update user when updating user (persistence): %v\n", err)
		return err
	}
	return nil
}

// 指定したidのユーザーを削除
func (p *UserPersistence) DeleteUser(id uint) error {
	return p.db.Delete(&model.User{}, id).Error
}
