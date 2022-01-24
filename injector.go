//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	logRepository "github.com/misyuari/go-omdb/domain/log/repository"
	logUseCase "github.com/misyuari/go-omdb/domain/log/use_case"
	"github.com/misyuari/go-omdb/domain/movie/delivery/rest"
	"github.com/misyuari/go-omdb/domain/movie/repository"
	"github.com/misyuari/go-omdb/domain/movie/use_case"
	"net/http"
)

var movieSet = wire.NewSet(
	repository.NewMovieRepository,
	use_case.NewMovieUseCase,
	rest.NewMovieHandler,
	rest.RegisterMovieRouter,
	logUseCase.NewLogUseCase,
	logRepository.NewLogRepository,
)

func InitializedHttpServer() *http.Server {
	wire.Build(
		movieSet,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		NewServer,
	)
	return nil
}
