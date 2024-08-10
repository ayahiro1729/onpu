package service

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/infrastructure/repository"
	"github.com/ayahiro1729/onpu/api/infrastructure/persistence"
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
