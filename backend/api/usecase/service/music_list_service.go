package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayahiro1729/onpu/api/domain/model"
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

func (s *MusicListService) GetTopTracks(accessToken string) ([]repository.MusicDTO, error) {
	uri := "https://api.spotify.com/v1/me/top/tracks?limit=10"
	req, err := http.NewRequest("GET", uri, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get top tracks: %v", resp.Status)
	}

	var result struct {
		Items []struct {
			Name  string `json:"name"`
			Album struct {
				Images []struct {
					URL string `json:"url"`
				} `json:"images"`
			} `json:"album"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			ExternalURLs struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	var tracks []repository.MusicDTO
	for _, item := range result.Items {
		track := repository.MusicDTO{
			Name:        item.Name,
			Image:       item.Album.Images[0].URL,
			ArtistName:  item.Artists[0].Name,
			SpotifyLink: item.ExternalURLs.Spotify,
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (s *MusicListService) CreateMusicList(userID int) (*model.MusicList, error) {
	musicListID, err := s.musicListPersistence.CreateMusicList(userID)
	if err != nil {
		fmt.Printf("error creating music list (service): %v\n", err)
		return nil, err
	}
	musicList := model.MusicList{
		ID: musicListID,
	}
	return &musicList, nil
}

func (s *MusicListService) CreateMusics(musicListID int, tracks []repository.MusicDTO) error {
	err := s.musicListPersistence.CreateMusics(musicListID, tracks)
	if err != nil {
		fmt.Printf("error creating musics (service): %v\n", err)
		return err
	}
	return nil
}
