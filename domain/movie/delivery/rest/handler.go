package rest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	logUseCase "github.com/misyuari/go-omdb/domain/log/use_case"
	"github.com/misyuari/go-omdb/domain/movie/model"
	"github.com/misyuari/go-omdb/domain/movie/use_case"
	"github.com/misyuari/go-omdb/helper"
	"net/http"
)

type MovieHandler interface {
	Detail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	List(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type MovieHandlerImpl struct {
	MovieUseCase use_case.MovieUseCase
	LogUseCase   logUseCase.LogUseCase
}

func NewMovieHandler(movieUseCase use_case.MovieUseCase, logUseCase logUseCase.LogUseCase) MovieHandler {
	return &MovieHandlerImpl{
		MovieUseCase: movieUseCase,
		LogUseCase:   logUseCase,
	}
}

func (handler *MovieHandlerImpl) Detail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	movieId := params.ByName("movieId")
	go handler.LogUseCase.Create("REST", "DETAIL_MOVIE", map[string]string{"id": movieId})
	response := handler.MovieUseCase.Detail(movieId)
	response.Code = helper.HTTPStatusFromCode(response.Code)
	helper.WriteToResponseBody(writer, response)
}

func (handler *MovieHandlerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	movieRequest := &model.MovieRequestList{
		Page:    query.Get("page"),
		Keyword: query.Get("keyword"),
	}
	fmt.Println("---------------123", query.Get("page"), query.Get("keyword"))
	go handler.LogUseCase.Create("REST", "LIST_MOVIE", movieRequest)
	response := handler.MovieUseCase.List(movieRequest)
	response.Code = helper.HTTPStatusFromCode(response.Code)
	helper.WriteToResponseBody(writer, response)
}
