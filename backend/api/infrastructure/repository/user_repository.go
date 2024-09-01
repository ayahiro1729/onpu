package repository

import (
	"github.com/ayahiro1729/onpu/api/domain/model"
)

type UserRepository interface {
	FindUserByID(id uint) (*model.User, error)
	FindUserBySpotifyID(spotifyID string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}