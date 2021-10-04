package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"userCrud/app/logic/user"
	"userCrud/app/netInterfaces/interfaces"
	"userCrud/app/repository"
)

func main() {
	listener, err := net.Listen("tcp", os.Getenv("PORT"))

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	userRepo, err := repository.GetUserRepository()
	if err != nil {
		grpclog.Fatalf("repository init error: %v", err)
	}


	userCRUD := interfaces.UserServer{UseCases: user.NewUserUseCases(userRepo)}
	grpcServer := grpc.NewServer()

	interfaces.RegisterUserCrudServer(grpcServer, &userCRUD)

	if err := grpcServer.Serve(listener); err != nil {
		grpclog.Fatal("failed to serve: %s", err)
	}
}
