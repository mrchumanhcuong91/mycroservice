package main

import (
	"authService/handler"
	pb "authService/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("authservice"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAuthServiceHandler(srv.Server(), new(handler.AuthService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
