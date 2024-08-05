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

	// Relations
	User   User    `gorm:"foreignKey:UserID"`
	Musics []Music `gorm:"foreignKey:MusicListID"`
}
