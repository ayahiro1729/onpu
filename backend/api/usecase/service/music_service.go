package service

import (
	"domain/model"
	"domain/repository"
)

type MusicListUsecase interface {
	LatestMusicList() (*repository.MusicListWithMusicDTO, error)
}

type musicListUsecase struct {
	mlr repository.MusicListRepository
}

func NewMusicListUsecase(mlr repository.MusicListRepository) MusicListUsecase {
	return &musicListUsecase{mlr}
}

func (mlu *musicListUsecase) LatestMusicList() (*repository.MusicListWithMusicDTO, error) {
	musicList, err := mlu.mlr.LatestMusicList()
	if err != nil {
		return nil, err
	}
	return musicList, nil
}