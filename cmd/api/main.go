package main

import (
	"context"
	"crud-grpc-gofiber/internal/routes"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()
	conn, err := grpc.Dial("localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	client := userpb.NewUserServiceClient(conn)

	ctx := context.Background()

	routes.SetupRouter(app, client, ctx)
	app.Listen(":3000")
	log.Println("goFiber service running on - http://localhost:3000/")
}
