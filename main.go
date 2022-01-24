package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/misyuari/go-omdb/config"
	MovieGrpc "github.com/misyuari/go-omdb/domain/movie/delivery/grpc"
	"github.com/misyuari/go-omdb/helper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    config.GlobalConfig.Env.RestHost,
		Handler: handler,
	}
}

func main() {
	config.LoadConfig()
	config.NewDB()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		httpServer := InitializedHttpServer()
		errHttp := httpServer.ListenAndServe()
		helper.PanicIfError(errHttp)
	}()
	go func() {
		defer waitGroup.Done()
		srv := grpc.NewServer()
		MovieGrpc.RegisterMoviesServer(srv)
		l, errGrpc := net.Listen("tcp", config.GlobalConfig.Env.GrpcHost)
		helper.PanicIfError(errGrpc)
		log.Fatal(srv.Serve(l))
	}()

	waitGroup.Wait()
}
