package apiservice

import (
	models "crud-grpc-gofiber/internal/models"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"log"
	"strconv"

	"crud-grpc-gofiber/internal/grpcclient"

	"github.com/gofiber/fiber/v2"
)

func ServerStatus(ctx *fiber.Ctx) error {
	res := models.GetApiResponse("api.server.status", "OK", "Server is Alive..! Ready with Service...")
	return ctx.JSON(res)
}

func GetUserHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	resp, err := grpcclient.Client.GetUser(grpcclient.Ctx, &userpb.GetUserRequest{Id: int32(id)})
	if err != nil {
		log.Printf("could not get user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	log.Println("gRPC Client got user info for user: ", resp.FullName)
	final_response := models.GetApiResponse("api.get.user", "OK", resp)
	return c.JSON(final_response)
}
