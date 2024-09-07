package model

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	ID         int       `gorm:"primaryKey";column:id`
	FollowerID int       `gorm:"column:follower_id;not null"`
	FolloweeID int       `gorm:"column:followee_id;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt  *time.Time `gorm:"column:deleted_at;autoDeleteTime"`

	// Relations
	Follower User `gorm:"foreignKey:FollowerID;references:ID"`
	Followee User `gorm:"foreignKey:FolloweeID;references:ID"`
}

func (Follow) TableName() string {
	return "follows"
}
