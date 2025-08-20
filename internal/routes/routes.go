package routes

import (
	"crud-grpc-gofiber/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", service.ServerStatus)
	v1.Get("/getUser/:id", service.GetUserHandler)
}
