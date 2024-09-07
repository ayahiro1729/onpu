package repository

import (
	"github.com/ayahiro1729/onpu/api/domain/model"
)

type UserSearchResultDTO struct {
	UserID				int			`json:"user_id"`
	UserName			string	`json:"user_name"`
	DisplayName		string	`json:"display_name"`
	IconImage			string	`json:"icon_image"`
}

type UserRepository interface {
	FindUserByID(id uint) (*model.User, error)
	FindUserBySpotifyID(spotifyID string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}