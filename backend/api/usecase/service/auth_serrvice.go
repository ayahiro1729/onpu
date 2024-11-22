package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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

	auth := base64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + cfg.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %w", err)
		}
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
func (s *AuthService) FetchSpotifyUser(token string) (*model.User, error) {
	uri := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", uri, http.NoBody)
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user profile. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
	}

	var spotifyUser struct {
		ID              string `json:"id"`
		DisplayName     string `json:"display_name"`
		Email           string `json:"email"`
		Country         string `json:"country"`
		ExplicitContent struct {
			FilterEnabled bool `json:"filter_enabled"`
			FilterLocked  bool `json:"filter_locked"`
		} `json:"explicit_content"`
		ExternalURL struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  interface{} `json:"href"`
			Total int         `json:"total"`
		} `json:"followers"`
		Href   string `json:"href"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
		Product string `json:"product"`
		Type    string `json:"type"`
	}

	if err := json.Unmarshal(bodyBytes, &spotifyUser); err != nil {
		return nil, fmt.Errorf("failed to decode user profile: %w", err)
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

func (s *AuthService) FetchUserID(user *model.User) (user_id uint, err error) {
	uri := "http://localhost:8080/api/v1/user"
	client := &http.Client{}
	jsonData, err := json.Marshal(user)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal user: %w", err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, fmt.Errorf("failed to read response body: %w", err)
		}
		return 0, fmt.Errorf("failed to create user. Status: %d, Body: %s", resp.StatusCode, string(body))
	}

	var respBody struct {
		UserID uint `json:"user_id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return 0, fmt.Errorf("failed to decode user_id: %w", err)
	}

	return respBody.UserID, nil
}
