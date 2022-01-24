package helper

import (
	"encoding/json"
	"fmt"
	"github.com/misyuari/go-omdb/config"
	"github.com/misyuari/go-omdb/domain/movie/model"
	"github.com/misyuari/go-omdb/helper"
)

type omdbMovieList struct {
	Search       *[]model.Movie `json:"Search"`
	TotalResults string         `json:"totalResults"`
	Response     string         `json:"Response"`
}

func FetchMovieDetail(id string) (*model.Movie, error) {
	URL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", config.GlobalConfig.Env.OmdbApiKey, id)

	response, err := helper.RequestGet(URL)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	movie := &model.Movie{}
	errDecode := json.NewDecoder(response.Body).Decode(movie)
	if errDecode != nil {
		return nil, errDecode
	}

	return movie, nil
}

func FetchMovieList(page uint32, keyword string) (*[]model.Movie, error) {
	URL := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", config.GlobalConfig.Env.OmdbApiKey, keyword, page)

	response, err := helper.RequestGet(URL)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	listMovie := &omdbMovieList{}
	errDecode := json.NewDecoder(response.Body).Decode(listMovie)
	if errDecode != nil {
		return nil, errDecode
	}

	return listMovie.Search, nil
}
