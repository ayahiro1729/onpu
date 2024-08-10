package repository

import (
	"time"
)

type MusicDTO struct {
	MusicID     int    `json:"music_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	ArtistName  string `json:"artist_name"`
	SpotifyLink string `json:"spotify_link"`
}

type MusicListDTO struct {
	MusicListID int       `json:"music_list_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type MusicListWithMusicDTO struct {
	MusicListID int        `json:"music_list_id"`
	CreatedAt   time.Time  `json:"created_at"`
	Musics      []MusicDTO `json:"musics"`
}

type MusicListRepository interface {
	LatestMusicList(userID int) (*MusicListWithMusicDTO, error)
}
