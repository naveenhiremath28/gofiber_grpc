package routes

import (
	"context"
	service "crud-grpc-gofiber/internal/service"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, client userpb.UserServiceClient, ctx context.Context) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", service.ServerStatus)
	v1.Get("/getUser", service.GetUserHandler(client, ctx))
}
