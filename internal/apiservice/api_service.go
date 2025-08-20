package apiservice

import (
	models "crud-grpc-gofiber/internal/models"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"encoding/json"
	"fmt"
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

func AddUserHandler(c *fiber.Ctx) error {
	request := new(models.ApiRequest)
	if err := c.BodyParser(request); err != nil {
		fmt.Println("error: ", err)
		res := models.GetApiResponse("api.add", "ERROR", c.Status(400).JSON(fiber.Map{"error": err.Error()}))
		return c.JSON(res)
	}
	var user models.User
	if err := json.Unmarshal(request.Request, &user); err != nil {
		res := models.GetApiResponse("api.add", "ERROR", c.Status(400).JSON(fiber.Map{"error": err.Error()}))
		return c.JSON(res)
	}
	fmt.Println("user: ", user)

	resp, err := grpcclient.Client.AddUser(grpcclient.Ctx, &userpb.AddUserRequest{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	})
	if err != nil {
		log.Printf("could not get user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	log.Println("gRPC Client added user: ", user.Username)
	final_response := models.GetApiResponse("api.get.user", "OK", resp)
	return c.JSON(final_response)
}
