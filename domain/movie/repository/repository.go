package repository

import (
	"github.com/misyuari/go-omdb/domain/movie/helper"
	"github.com/misyuari/go-omdb/domain/movie/model"
)

type MovieRepository interface {
	Detail(id string) (*model.Movie, error)
	List(page uint32, keyword string) (*[]model.Movie, error)
}

type MovieRepositoryImpl struct{}

func (*MovieRepositoryImpl) Detail(id string) (*model.Movie, error) {
	movie, err := helper.FetchMovieDetail(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (*MovieRepositoryImpl) List(page uint32, keyword string) (*[]model.Movie, error) {
	listMovie, err := helper.FetchMovieList(page, keyword)
	if err != nil {
		return nil, err
	}
	return listMovie, nil
}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}
