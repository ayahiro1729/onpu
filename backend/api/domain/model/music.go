package model

import (
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

	// Relations
	MusicList MusicList `gorm:"foreignKey:MusicListID"`
}
