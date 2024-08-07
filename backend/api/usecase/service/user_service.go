package service

import (
    "errors"

    "github.com/ayahiro1729/onpu/api/domain/model"
    "myapp/domain/repository"
)

type UserService struct {
    userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterOrLogin(spotifyUser *model.User) (*model.User, error) {
    // Spotify IDでユーザーを検索
    user, err := s.userRepository.FindBySpotifyID(spotifyUser.SpotifyID)
		// ユーザーが既に存在する場合、ログインと見なしてそのユーザーを返す
    if err == nil {
        return user, nil
    }

    // ユーザーが存在しない場合は新規登録
    newUser := &model.User{
        SpotifyID:     spotifyUser.ID,
        UserName:      spotifyUser.DisplayName,
        DisplayName:   spotifyUser.DisplayName,
				IconImage:     spotifyUser.IconImage,
				ThemeID:       0,
				XLink:         "",
				InstagramLink: "",
    }

    err = s.userRepository.CreateUser(newUser)
    if err != nil {
        return nil, errors.New("failed to create new user")
    }

    return newUser, nil
}
