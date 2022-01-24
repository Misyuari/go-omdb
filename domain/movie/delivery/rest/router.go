package rest

import (
	"github.com/julienschmidt/httprouter"
	"github.com/misyuari/go-omdb/helper"
)

func RegisterMovieRouter(handler MovieHandler) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/movie/:movieId", handler.Detail)
	router.GET("/api/movie", handler.List)

	router.PanicHandler = helper.PanicHandler
	return router
}
