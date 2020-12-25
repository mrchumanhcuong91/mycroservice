package main

import (
	"userService/handler"
	pb "userService/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("userservice"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), new(handler.UserService))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
