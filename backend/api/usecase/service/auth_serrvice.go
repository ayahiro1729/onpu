package service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ayahiro1729/onpu/api/config"
	"github.com/ayahiro1729/onpu/api/domain/model"
)

type AuthService struct {
	spotifyConfig *config.SpotifyConfig
}

func NewAuthService(spotifyConfig *config.SpotifyConfig) *AuthService {
	return &AuthService{
		spotifyConfig: spotifyConfig,
	}
}

// Spotifyの認証コードを使用してアクセストークンを取得
func (s *AuthService) FetchSpotifyToken(code string) (string, error) {
	uri := "https://accounts.spotify.com/api/token"
	cfg, err := config.NewSpotifyConfig()
	if err != nil {
		return "", err
	}

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", cfg.RedirectURI)

	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte( cfg.ClientID + ":" + cfg.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
    return "", fmt.Errorf("failed to exchange authorization code for token. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
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
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
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
