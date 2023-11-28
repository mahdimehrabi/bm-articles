package grpc

import (
	"bm/src/apps/grpc/proto/article"
	article2 "bm/src/apps/grpc/servers/article"
	"bm/src/domain/article/repositories/gorm"
	"bm/src/domain/article/services"
	"crypto/tls"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}

func RunServer() {
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials ", err)
	}

	lis, err := net.Listen("tcp", "localhost:6001")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(tlsCredentials))

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
