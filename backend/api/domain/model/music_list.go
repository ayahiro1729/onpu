package model

import (
	"time"

	"gorm.io/gorm"
)

type MusicList struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time `gorm:"autoDeleteTime"`

	// Relations
	User   User    `gorm:"foreignKey:UserID"`
	Musics []Music `gorm:"foreignKey:MusicListID"`
}

func (MusicList) TableName() string {
	return "music_lists"
}
