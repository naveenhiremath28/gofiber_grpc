package main

import (
	"context"
	"crud-grpc-gofiber/internal/routes"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"log"

	"crud-grpc-gofiber/internal/grpcclient"

	"github.com/gofiber/fiber/v2"
)

var (
	Client userpb.UserServiceClient
	Ctx    context.Context
)

func main() {
	app := fiber.New()
	grpcclient.InitGRPCClient()

	routes.SetupRouter(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
