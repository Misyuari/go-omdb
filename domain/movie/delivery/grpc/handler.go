package grpc

import (
	"context"
	logRepository "github.com/misyuari/go-omdb/domain/log/repository"
	logUseCase "github.com/misyuari/go-omdb/domain/log/use_case"
	"github.com/misyuari/go-omdb/domain/movie/delivery/grpc/proto"
	"github.com/misyuari/go-omdb/domain/movie/model"
	"github.com/misyuari/go-omdb/domain/movie/repository"
	"github.com/misyuari/go-omdb/domain/movie/use_case"
	"github.com/misyuari/go-omdb/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type MoviesServer struct {
	MovieUseCase use_case.MovieUseCase
	LogUseCase   logUseCase.LogUseCase
}

func (server MoviesServer) Detail(ctx context.Context, request *proto.MovieRequestDetail) (*proto.MovieDetail, error) {
	go server.LogUseCase.Create("GRPC", "DETAIL_MOVIE", request)
	response := server.MovieUseCase.Detail(request.Id)
	if response.Code != codes.OK {
		return nil, status.Error(response.Code, response.Message)
	}
	movie := &proto.Movie{}
	err := helper.ConvertGrpcStruct(response.Data, movie)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	protoMovie := &proto.MovieDetail{
		Code:    uint32(response.Code),
		Message: response.Message,
		Data:    movie,
	}

	return protoMovie, nil
}

func (server MoviesServer) List(ctx context.Context, request *proto.MovieRequestList) (*proto.MovieList, error) {
	movieRequest := &model.MovieRequestList{
		Page:    strconv.Itoa(int(request.Page)),
		Keyword: request.Keyword,
	}
	go server.LogUseCase.Create("GRPC", "LIST_MOVIE", movieRequest)
	response := server.MovieUseCase.List(movieRequest)
	if response.Code != codes.OK {
		return nil, status.Error(response.Code, response.Message)
	}
	movies := &[]*proto.Movie{}
	err := helper.ConvertGrpcStruct(response.Data, movies)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	protoMovies := &proto.MovieList{
		Code:    uint32(response.Code),
		Message: response.Message,
		Data:    *movies,
	}
	return protoMovies, nil
}

func RegisterMoviesServer(grpcServer *grpc.Server) {
	var movieServer = MoviesServer{
		MovieUseCase: use_case.NewMovieUseCase(repository.NewMovieRepository()),
		LogUseCase:   logUseCase.NewLogUseCase(logRepository.NewLogRepository()),
	}
	proto.RegisterMoviesServer(grpcServer, movieServer)
}
