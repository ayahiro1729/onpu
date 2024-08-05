package model

type Music struct {
	ID int `gorm:"primaryKey"`
	MusicListID int `gorm:"not null"`
	Name         string `gorm:"not null"`
	Image        string
	ArtistName   string
	SpotifyLink  string

	// Relations
	MusicList MusicList `gorm:"foreignKey:MusicListID"`
}
