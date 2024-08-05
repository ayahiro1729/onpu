package model

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	ID         int       `gorm:"primaryKey"`
	FollowerID int       `gorm:"not null"`
	FolloweeID int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	// Relations
	Follower User `gorm:"foreignKey:FollowerID"`
	Followee User `gorm:"foreignKey:FolloweeID"`
}
