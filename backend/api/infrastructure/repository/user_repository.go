package repository

import (
	"github.com/ayahiro1729/onpu/api/domain/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// id でユーザーを検索
func (r *UserRepository) FindUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// spotify_idでユーザーを検索
func (r *UserRepository) FindUserBySpotifyID(spotifyID string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("spotify_id = ?", spotifyID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 新しいユーザーを作成
func (r *UserRepository) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーを更新
func (r *UserRepository) UpdateUser(user *model.User) error {
	// ユーザーが存在するか確認
	var existingUser model.User
	if err := r.db.First(&existingUser, user.ID).Error; err != nil {
		return err
	}

	// レコードを更新
	if err := r.db.Model(existingUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// 指定したidのユーザーを削除
func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
