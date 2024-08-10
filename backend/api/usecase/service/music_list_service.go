package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ayahiro1729/onpu/api/domain/types"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
)

// type MusicListService interface {
// 	LatestMusicList(userID int) (*repository.MusicListWithMusicDTO, error)
// }

type MusicListService struct {
	musicListPersistence persistence.MusicListPersistence
}

func NewMusicListService(musicListPersistence persistence.MusicListPersistence) *MusicListService {
	return &MusicListService{musicListPersistence: musicListPersistence}
}

func (s *MusicListService) LatestMusicList(userID int) (*repository.MusicListWithMusicDTO, error) {
	musicList, err := s.musicListPersistence.LatestMusicList(userID)
	if err != nil {
		fmt.Printf("error getting latest music list (service): %v\n", err)
		return nil, err
	}
	return musicList, nil
}

// セッションに保存されたアクセストークンを取得
func (s *MusicListService) CheckAccessToken() (string, error) {
	uri := "http://localhost:8080/api/v1/session/token"
	req, err := http.NewRequest("GET", uri, http.NoBody)
	if err != nil {
		fmt.Printf("error creating token request: %v\n", err)
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sending token request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading token response: %v\n", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to read response body. Status: %d, Body: %s", resp.StatusCode, bodyBytes)
	}

	var respBody struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		fmt.Printf("error decoding token response: %v\n", err)
		return "", err
	}

	return respBody.AccessToken, nil
}

// users.idを取得
func (s *MusicListService) FetchUserID(token string) (int, error) {
	uri := "http://localhost:8080/api/v1/user"
	req, err := http.NewRequest("GET", uri, http.NoBody)
	if err != nil {
		fmt.Printf("error creating user request: %v\n", err)
		return 0, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sending user request: %v\n", err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error getting user ID: %v\n", err)
		return 0, err
	}

	var respBody struct {
		UserID int `json:"user_id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		fmt.Printf("error decoding user response: %v\n", err)
		return 0, err
	}

	return respBody.UserID, nil
}

// ユーザーのお気に入りの曲を10曲取得
func (s *MusicListService) FetchTenFavoriteMusics(token string) ([]types.TrackItem, error) {
	uri := "https://api.spotify.com/v1/me/top/tracks?limit=10"
	req, err := http.NewRequest("GET", uri, http.NoBody)
	if err != nil {
		fmt.Printf("error creating request: %v\n", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error sending request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user profile. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
	}

	var respBody struct {
		Items []struct {
			Name   string `json:"name"`
			Images []struct {
				URL string `json:"url"`
			} `json:"images"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			ExternalURL struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		fmt.Printf("error decoding response: %v\n", err)
		return nil, err
	}

	tracks := []types.TrackItem{}
	for _, item := range respBody.Items {
		tracks = append(tracks, types.TrackItem{
			Name:        item.Name,
			Images:      item.Images,
			Artists:     item.Artists,
			ExternalURL: item.ExternalURL,
		})
	}

	return tracks, nil
}

// music_listのデータを作成
func (s *MusicListService) CreateMusicList(userID int) error {
	s.musicListPersistence.SaveMusicList(userID)
	return nil
}

// ユーザーの最新のmusic_listのidを取得
func (s *MusicListService) GetLatestMusicListID(userID int) (int, error) {
	musicListID, err := s.musicListPersistence.LatestMusicListID(userID)
	if err != nil {
		fmt.Printf("error getting latest music list (service): %v\n", err)
		return 0, err
	}
	return musicListID, nil
}

// musicのデータを作成
func (s *MusicListService) CreateSingleMusics(musicListId int, musics []types.TrackItem) error {
	s.musicListPersistence.SaveMusics(musicListId, musics)
	return nil
}
