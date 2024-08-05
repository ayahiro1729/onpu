package model

import (
	"time"
)

type MusicList struct {
	ID int `gorm:"primaryKey"`
	UserID int `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Relations
	User User `gorm:"foreignKey:UserID"`
	Musics []Music `gorm:"foreignKey:MusicListID"`
}
