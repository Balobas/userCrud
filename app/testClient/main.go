package main

import (
	"google.golang.org/grpc"
	"log"
	"os"
	"userCrud/app/netInterfaces/interfaces"
	"userCrud/app/testClient/tests"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := interfaces.NewUserCrudClient(conn)

	tests.NewUserCrudServerTester(c).Run()
}
