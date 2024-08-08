package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	spotifyConfig *config.SpotifyConfig
}

func NewAuthService(spotifyConfig *config.SpotifyConfig) *AuthService {
	return &AuthService{
		spotifyConfig: spotifyConfig,
	}
}

func (s *AuthService) GetSpotifyAuthURL() string {
	return "https://accounts.spotify.com/authorize?client_id=" + s.spotifyConfig.ClientID +
		"&response_type=code&redirect_uri=" + s.spotifyConfig.RedirectURI +
		"&scope=user-read-private user-read-email"
}

// Spotifyの認証コードを使用してアクセストークンを取得
func (s *AuthService) GetSpotifyToken(code string) (string, error) {
	uri := "https://accounts.spotify.com/api/token"

	reqBody := map[string]string{
		"grant_type":    "authorization_code",
		"code":          code,
		"redirect_uri":  s.spotifyConfig.RedirectURI,
		"client_id":     s.spotifyConfig.ClientID,
		"client_secret": s.spotifyConfig.ClientSecret,
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(uri, "application/x-www-form-urlencoded", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to exchange authorization code for token")
	}

	var respBody struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}

	return respBody.AccessToken, nil
}

// Spotifyのアクセストークンを使用してユーザ情報を取得
func (s *AuthService) GetSpotifyUser(token string) (*model.User, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user profile")
	}

	var spotifyUser struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
		Images      []struct {
			URL string `json:"url"`
		} `json:"images"`
	}

	user := &model.User{
		SpotifyID:     spotifyUser.ID,
		UserName:      spotifyUser.DisplayName,
		DisplayName:   spotifyUser.DisplayName,
		IconImage:     "",
		ThemeID:       0,
		XLink:         "",
		InstagramLink: "",
	}
	if len(spotifyUser.Images) > 0 {
		user.IconImage = spotifyUser.Images[0].URL
	}

	return user, nil
}

// Spotifyのユーザ情報を使用してJWTトークンを生成
func (s *AuthService) GenerateJWTToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.SpotifyID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.spotifyConfig.JWTSecret))
}
