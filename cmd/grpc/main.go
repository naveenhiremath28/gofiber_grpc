package main

import (
	"log"
	"net"

	"crud-grpc-gofiber/internal/database"
	"crud-grpc-gofiber/internal/database/dbmodels"
	grcpservice "crud-grpc-gofiber/internal/grpcservice"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"

	"google.golang.org/grpc"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&dbmodels.User{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &grcpservice.UserServer{})

	log.Println("Server listening to port: http://localhost:8080/")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
