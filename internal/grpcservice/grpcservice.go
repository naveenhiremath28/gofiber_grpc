package grcpservice

import (
	"context"
	"crud-grpc-gofiber/internal/database/dbservice"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
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
	log.Printf("Received AddUser request: username=%s, email=%s, fullname=%s",
		req.Username, req.Email, req.FullName)

	result := dbservice.AddEmployee(req)
	res := "Failed to Insert Record...!"
	if result {
		res = "Successfully Inserted...!"
	}

	return &userpb.AddUserResponse{
		Status: res,
	}, nil
}
