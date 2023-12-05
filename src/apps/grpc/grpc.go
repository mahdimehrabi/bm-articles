package grpc

import (
	article2 "bm/src/apps/grpc/servers/article"
	"bm/src/domain/article/repositories/gorm"
	"bm/src/domain/article/services"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/mahdimehrabi/bm-articles/src/apps/grpc/proto/article"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	//Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
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
