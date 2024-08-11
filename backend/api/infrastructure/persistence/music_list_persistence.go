package persistence

import (
	"fmt"

	"github.com/ayahiro1729/onpu/api/domain/model"
	"github.com/ayahiro1729/onpu/api/domain/types"
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

	if err := mlp.db.Model(&model.MusicList{}).
		Select("id AS music_list_id", "created_at").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(1).
		Scan(&musicList).Error; err != nil {
		fmt.Printf("error during select from music_list (persistence): %v\n", err)
		return nil, err
	}

	musics := []repository.MusicDTO{}

	if err := mlp.db.Model(&model.Music{}).
		Select("id AS music_id", "name", "image", "artist_name", "spotify_link").
		Where("music_list_id = ?", musicList.MusicListID).
		Scan(&musics).Error; err != nil {
		fmt.Printf("error during select from music (persistence): %v\n", err)
		return nil, err
	}

	return &repository.MusicListWithMusicDTO{
		MusicListID: musicList.MusicListID,
		CreatedAt:   musicList.CreatedAt,
		Musics:      musics,
	}, nil
}

// music_listをDBに保存
func (mlp *MusicListPersistence) SaveMusicList(userID int) error {
	fmt.Printf("Saving music list for user ID: %d\n", userID)
	if err := mlp.db.Create(&model.MusicList{
		UserID: userID,
	}).Error; err != nil {
		fmt.Printf("error during insert into music_list (persistence): %v\n", err)
		return err
	}

	return nil
}

// ユーザーの最新のmusic_listのidを取得
func (mlp *MusicListPersistence) LatestMusicListID(userID int) (int, error) {
	musicList := model.MusicList{}

	if err := mlp.db.Model(&model.MusicList{}).
		Select("id").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(1).
		Scan(&musicList).Error; err != nil {
		fmt.Printf("error during select from music_list (persistence): %v\n", err)
		return 0, err
	}

	return musicList.ID, nil
}

// 10件のmusicをDBに保存
func (mlp *MusicListPersistence) SaveMusics(musicListId int, musics []types.TrackItem) error {
	for _, music := range musics {
		if err := mlp.db.Create(&model.Music{
			MusicListID: musicListId,
			Name:        music.Name,
			Image:       music.Images[0].URL,
			ArtistName:  music.Artists[0].Name,
			SpotifyLink: music.ExternalURL.Spotify,
		}).Error; err != nil {
			fmt.Printf("error during insert into music (persistence): %v\n", err)
			return err
		}
	}

	return nil
}
