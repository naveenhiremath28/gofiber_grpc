package grpcclient

import (
	"context"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	Client userpb.UserServiceClient
	Ctx    context.Context
)

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	Client = userpb.NewUserServiceClient(conn)
	Ctx = context.Background()
}