package model

import (
	"time"

	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	MusicListID int    `gorm:"not null"`
	Name        string `gorm:"not null"`
	Image       string
	ArtistName  string
	SpotifyLink string
	CreatedAt   *time.Time `gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime"`
	DeletedAt   *time.Time `gorm:"autoDeleteTime"`

	// Relations
	MusicList MusicList `gorm:"foreignKey:MusicListID"`
}

func (Music) TableName() string {
	return "musics"
}
