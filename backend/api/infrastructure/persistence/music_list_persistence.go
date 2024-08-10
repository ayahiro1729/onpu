package persistence

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/infrastructure/repository"

	"gorm.io/gorm"
)

type MusicListPersistence struct {
	db *gorm.DB
}

func NewMusicListPersistence(db *gorm.DB) *MusicListPersistence {
	return &MusicListPersistence{db: db}
}

func (mlp *MusicListPersistence) LatestMusicList(userID int) (*repository.MusicListWithMusicDTO, error) {
	musicList := repository.MusicListDTO{}

	errMusicList := mlp.db.Model(&model.MusicList{}).
			Select("id AS music_list_id", "created_at").
			Where("user_id = ?", userID).
			Order("created_at desc").
			Limit(1).
			Scan(&musicList).Error

	if errMusicList != nil {
		fmt.Printf("error during select from music_list (persistence): %v\n", errMusicList)
		return nil, errMusicList
	}

	musics := []repository.MusicDTO{}

	errMusic := mlp.db.Model(&model.Music{}).
			Select("id AS music_id", "name", "image", "artist_name", "spotify_link").
			Where("music_list_id = ?", musicList.MusicListID).
			Scan(&musics).Error
	
	if errMusic != nil {
		fmt.Printf("error during select from music (persistence): %v\n", errMusic)
		return nil, errMusic
	}

	return &repository.MusicListWithMusicDTO{
		MusicListID: musicList.MusicListID,
		CreatedAt: musicList.CreatedAt,
		Musics: musics,
	}, nil
}
