package grcpservice

import (
	"context"
	"crud-grpc-gofiber/internal/database/dbservice"
	"crud-grpc-gofiber/internal/models"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"fmt"
	"log"
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

func (s *UserServer) AddUser(ctx context.Context, req *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	fmt.Println("\n\n\nreq: ", req)
	log.Printf("\n\n\nReceived AddUser request: username=%s, email=%s, fullname=%s",
		req.Username, req.Email, req.FullName)

	result := dbservice.AddEmployee(req)
	res := "Failed to Insert Record...!"
	if result {
		res = "Successfully Inserted Record...!"
	}

	return &userpb.AddUserResponse{
		Status: res,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	log.Println("gRPC Server deleting user info for Id: ", req.Id)

	result := dbservice.DeleteUser(req.Id)
	res := "Failed to Delete Record...!"
	if result {
		res = "Successfully Deleted Record...!"
	}

	return &userpb.DeleteUserResponse{
		Status: res,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	fmt.Println("\n\n\nreq: ", req)
	log.Printf("\n\n\nReceived AddUser request: username=%s, email=%s, fullname=%s",
		req.Username, req.Email, req.FullName)
	var user models.User
	user.ID = int(req.Id)
	user.Username = req.Username
	user.Email = req.Email
	user.FullName = req.FullName
	result := dbservice.UpdateUser(user)
	res := "Failed to Insert Record...!"
	if result {
		res = "Successfully Inserted Record...!"
	}

	return &userpb.UpdateUserResponse{
		Status: res,
	}, nil
}
