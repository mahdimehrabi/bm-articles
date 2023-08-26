package main

import (
	"bm/src/apps/gin"
	"bm/src/apps/grpc"
)

func main() {
	go gin.RunServer()
	grpc.RunServer()
}
