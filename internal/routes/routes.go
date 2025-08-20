package routes

import (
	"crud-grpc-gofiber/internal/apiservice"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", apiservice.ServerStatus)
	v1.Get("/getUser/:id", apiservice.GetUserHandler)
	v1.Post("/addUser", apiservice.AddUserHandler)
	v1.Delete("/deleteUser/:id", apiservice.DeleteUserHandler)
}
