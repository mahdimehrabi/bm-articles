package grpc

import (
	"bm/src/apps/grpc/proto/article"
	article2 "bm/src/apps/grpc/servers/article"
	"bm/src/domain/article/repositories/gorm"
	"bm/src/domain/article/services"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RunServer() {
	lis, err := net.Listen("tcp", "localhost:6001")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	articleRepo, err := gorm.NewGormArticleRepository()
	if err != nil {
		log.Fatal(err.Error())
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err.Error())
	}
	articleService := services.NewArticleService(articleRepo, logger)
	articleServer := article2.NewArticleServer(articleService)
	article.RegisterArticleServer(grpcServer, articleServer)
	reflection.Register(grpcServer)
	fmt.Println("listening grpc server at port 6001")
	grpcServer.Serve(lis)
}
