package main

import (
	"context"
	"log"
	"net"

	"crud-grpc-gofiber/internal/database"
	"crud-grpc-gofiber/internal/database/dbmodels"
	"crud-grpc-gofiber/internal/database/dbservice"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"

	"google.golang.org/grpc"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	log.Println("gRPC Server getting user info for Id: ", req.Id)

	user, err := dbservice.GetUser(req.Id)
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	return &userpb.GetUserResponse{
		Username:  user[0].Username,
		Email:     user[0].Email,
		FullName:  user[0].FullName,
		CreatedAt: user[0].CreatedAt.String(),
		UpdatedAt: user[0].UpdatedAt.String(),
	}, nil
}

func main() {
	database.Connect()
	database.DB.AutoMigrate(&dbmodels.User{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &UserServer{})

	log.Println("Server listening to port: http://localhost:8080/")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
