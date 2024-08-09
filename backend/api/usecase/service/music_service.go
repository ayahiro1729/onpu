package service

import (
	"domain/model"
	"domain/repository"
)

type MusicUsecase interface {
	LatestMusicList() 
}