package service

import (
	"errors"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
)

type UserService struct {
	userPersistence persistence.UserPersistence
}

func NewUserService(userPersistence persistence.UserPersistence) *UserService {
	return &UserService{userPersistence: userPersistence}
}

func (s *UserService) RegisterOrLogin(spotifyUser *model.User) (*model.User, error) {
	// Spotify IDでユーザーを検索
	user, err := s.userPersistence.FindUserBySpotifyID(spotifyUser.SpotifyID)
	// ユーザーが既に存在する場合、ログインと見なしてそのユーザーを返す
	if err == nil {
		return user, nil
	}

	// ユーザーが存在しない場合は新規登録
	newUser := &model.User{
		SpotifyID:     spotifyUser.SpotifyID,
		UserName:      spotifyUser.DisplayName,
		DisplayName:   spotifyUser.DisplayName,
		IconImage:     spotifyUser.IconImage,
		ThemeID:       0,
		XLink:         "",
		InstagramLink: "",
	}

	err = s.userPersistence.CreateUser(newUser)
	if err != nil {
		return nil, errors.New("failed to create new user")
	}

	return newUser, nil
}

func (s *UserService) FindUserProfile(id uint) (*model.User, error) {
	user, err := s.userPersistence.FindUserByID(id)
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	return user, nil
}

func (s *UserService) UpdateUserProfile(user *model.User) error {
	err := s.userPersistence.UpdateUser(user)
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (s *UserService) SearchUsers(search_string string) (*[]repository.UserSearchResultDTO, error) {
	users, err := s.userPersistence.FindUsersByUserName(search_string)
	if err != nil {
		return nil, errors.New("failed to find user by user_name")
	}

	// usersが空でも空でなくてもそのまま返す
	return users, nil
}