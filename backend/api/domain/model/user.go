package model

import (
	"time"
)

type User struct {
	ID int `gorm:"primaryKey"`
	SpotifyID string `gorm:"unique;not null"`
	UserName string `gorm:"not null"`
	DisplayName string
	IconImage string
	ThemeID int
	XLink string
	InstagramLink string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relations
	MusicLists []MusicList `gorm:"foreignKey:UserID"`
	Followers []Follow `gorm:"foreignKey:FolloweeID"`
	Followees []Follow `gorm:"foreignKey:FollwerID"`
}
