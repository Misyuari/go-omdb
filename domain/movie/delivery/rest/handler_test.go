package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/misyuari/go-omdb/domain/movie/model"
	"github.com/misyuari/go-omdb/helper"
	"github.com/misyuari/go-omdb/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovieHandlerDetail(t *testing.T) {
	responseData := map[string]*helper.Response{
		"tt0372784": {
			Code:    codes.OK,
			Message: "Detail Movie",
			Data:    nil,
		},
		"tt0372711": {
			Code:    codes.NotFound,
			Message: "Movie Not Found",
			Data:    nil,
		},
	}
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
		params  httprouter.Params
	}
	request := httptest.NewRequest("GET", "/api/movie/tt0372784", nil)

	mockMovieUsecase := new(mocks.MovieUseCase)
	mockLogUsecase := new(mocks.LogUseCase)
	handler := &MovieHandlerImpl{
		MovieUseCase: mockMovieUsecase,
		LogUseCase:   mockLogUsecase,
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should success get detail movie",
			args: args{
				writer:  httptest.NewRecorder(),
				request: request,
				params:  httprouter.Params{{Key: "movieId", Value: "tt0372784"}},
			},
		},
		{
			name: "Should error not found",
			args: args{
				writer:  httptest.NewRecorder(),
				request: request,
				params:  httprouter.Params{{Key: "movieId", Value: "tt0372711"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movieId := tt.args.params.ByName("movieId")
			mockMovieUsecase.On("Detail", movieId).Return(responseData[movieId])
			mockLogUsecase.On("Create", "REST", "DETAIL_MOVIE", map[string]string{"id": movieId}).Return(responseData[movieId])

			handler.Detail(tt.args.writer, tt.args.request, tt.args.params)
			response := tt.args.writer.(*httptest.ResponseRecorder).Result()
			body, _ := io.ReadAll(response.Body)
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			assert.Equal(t, responseData[movieId].Message, result["message"])
		})
	}
}

func TestMovieHandlerList(t *testing.T) {
	responseData := &helper.Response{
		Code:    codes.OK,
		Message: "List Movie",
		Data:    nil,
	}
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
		params  httprouter.Params
	}
	request := httptest.NewRequest("GET", "/api/movie", nil)
	query := request.URL.Query()
	query.Add("page", "1")
	query.Add("keyword", "batman")
	request.URL.RawQuery = query.Encode()
	mockMovieUsecase := new(mocks.MovieUseCase)
	mockLogUsecase := new(mocks.LogUseCase)
	handler := &MovieHandlerImpl{
		MovieUseCase: mockMovieUsecase,
		LogUseCase:   mockLogUsecase,
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should success get detail movie",
			args: args{
				writer:  httptest.NewRecorder(),
				request: request,
				params:  httprouter.Params{{Key: "movieId", Value: "tt0372784"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			movieId := tt.args.params.ByName("movieId")
			movieRequest := &model.MovieRequestList{
				Page:    "1",
				Keyword: "batman",
			}
			mockMovieUsecase.On("List", movieRequest).Return(responseData)
			mockLogUsecase.On("Create", "REST", "DETAIL_MOVIE", map[string]string{"id": movieId}).Return(responseData)

			handler.List(tt.args.writer, tt.args.request, tt.args.params)
			response := tt.args.writer.(*httptest.ResponseRecorder).Result()
			body, _ := io.ReadAll(response.Body)
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			assert.Equal(t, responseData.Message, result["message"])
		})
	}
}
