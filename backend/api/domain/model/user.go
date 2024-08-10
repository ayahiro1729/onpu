package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int    `gorm:"primaryKey"`
	SpotifyID     string `gorm:"unique;not null"`
	UserName      string `gorm:"not null"`
	DisplayName   string
	IconImage     string
	ThemeID       int
	XLink         string
	InstagramLink string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     time.Time `gorm:"autoDeleteTime"`

	// Relations
	MusicLists []MusicList `gorm:"foreignKey:UserID"`
	Followers  []Follow    `gorm:"foreignKey:followee_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Following  []Follow    `gorm:"foreignKey:followee_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}
