package service

import (
	"context"
	models "crud-grpc-gofiber/internal/models"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ServerStatus(ctx *fiber.Ctx) error {
	res := models.GetApiResponse("api.server.status", "OK", "Server is Alive..! Ready with Service...")
	return ctx.JSON(res)
}

// func getUser(client userpb.UserServiceClient, ctx context.Context) error {
// 	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{Id: 1})
// 	if err != nil {
// 		log.Fatalf("could not get order: %v", err)
// 		return nil
// 	}
// 	fmt.Println("resp: ", resp)
// 	return resp
// }

func GetUserHandler(client userpb.UserServiceClient, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp, err := client.GetUser(ctx, &userpb.GetUserRequest{Id: 1})
		if err != nil {
			log.Printf("could not get user: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to fetch user",
			})
		}
		log.Println("gRPC Client got user info for user: ", resp.FullName)
		// Return the response as JSON
		return c.JSON(resp)
	}
}
