package persistence

import (
	"fmt"

	"domain/model"
	"domain/repository"

	"gorm.io/gorm"
)

import (
	"domain/model"
	"domain/repository"

	"gorm.io/gorm"
)

type MusicListPersistence struct {
	db *gorm.DB
}

func NewMusicListPersistence(db *gorm.DB) repository.MusicListRepository {
	return &MusicListPersistence{db}
}

func (mlp *MusicListPersistence) LatestMusicList() (repository.MusicListWithMusicDTO, error) {
	musicList := repository.MusicListDTO{}

	err := mlp.db.Model(&MusicList{}).
			Select("id AS music_list_id", "created_at").
			Where("user_id = ?", user_id).
			Order("created_at desc").
			Limit(1).
			Scan(&musicList).Error

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	musics := []repository.MusicDTO{}

	err := mlp.db.Model(&Musics{}).
			Select("id AS music_id", "name", "image", "artist_name", "spotify_link").
			Where("music_list_id = ?", musicList.MusicListID).
			Scan(&musics).Error
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
