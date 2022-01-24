package use_case

import (
	"github.com/misyuari/go-omdb/config"
	"github.com/misyuari/go-omdb/domain/movie/model"
	"github.com/misyuari/go-omdb/domain/movie/repository"
	"github.com/misyuari/go-omdb/helper"
	"google.golang.org/grpc/codes"
	"strconv"
)

type MovieUseCase interface {
	Detail(id string) *helper.Response
	List(request *model.MovieRequestList) *helper.Response
}

type MovieUseCaseImpl struct {
	repository repository.MovieRepository
}

func (useCase *MovieUseCaseImpl) Detail(id string) *helper.Response {
	movie, err := useCase.repository.Detail(id)
	if movie.Response == "False" || err != nil {
		return helper.WrapResponse(codes.NotFound, nil, "Movie Not Found")
	}
	return helper.WrapResponse(codes.OK, movie, "Detail Movie")
}

func (useCase *MovieUseCaseImpl) List(request *model.MovieRequestList) *helper.Response {
	err := config.GlobalConfig.Validator.Struct(request)
	if err != nil {
		return helper.WrapResponse(codes.FailedPrecondition, nil, err.Error())
	}
	page, _ := strconv.ParseUint(request.Page, 10, 32)
	movies, err := useCase.repository.List(uint32(page), request.Keyword)
	if movies == nil || err != nil {
		return helper.WrapResponse(codes.NotFound, nil, "Movie Not Found")
	}
	return helper.WrapResponse(codes.OK, movies, "List Movie")
}

func NewMovieUseCase(repository repository.MovieRepository) MovieUseCase {
	return &MovieUseCaseImpl{repository: repository}
}
