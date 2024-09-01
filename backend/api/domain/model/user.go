package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int    `gorm:"primaryKey" json:"id"`
	SpotifyID     string `gorm:"unique;not null" json:"spotify_id"`
	UserName      string `gorm:"not null" json:"user_name"`
	DisplayName   string `json:"display_name"`
	IconImage     string `json:"icon_image"`
	ThemeID       int `json:"theme_id"`
	XLink         string `json:"x_link"`
	InstagramLink string `json:"instagram_link"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     time.Time `gorm:"autoDeleteTime" json:"deleted_at"`

	// Relations
	MusicLists []MusicList `gorm:"foreignKey:UserID"`
	Followers  []Follow    `gorm:"foreignKey:followee_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Following  []Follow    `gorm:"foreignKey:followee_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}
