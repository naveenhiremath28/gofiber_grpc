package dbservice

import (
	"crud-grpc-gofiber/internal/database"
	"crud-grpc-gofiber/internal/models"
	userpb "crud-grpc-gofiber/pkg/protocolbuffers"
	"time"
)

func GetUser(id string) ([]models.User, error) {
	var user []models.User
	result := database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func AddEmployee(req *userpb.AddUserRequest) bool {
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		FullName:  req.FullName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return false
	}
	return true
}

func DeleteUser(id string) bool {
	var user []models.User
	result := database.DB.Delete(&user, "id = ?", id)
	if result.Error != nil {
		return false
	}
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userToUpdate models.User) bool {
	var user models.User
	result := database.DB.First(&user, "id = ?", userToUpdate.ID)
	if result.Error != nil {
		return false
	}
	user.Username = userToUpdate.Username
	user.FullName = userToUpdate.FullName
	user.Email = userToUpdate.Email
	user.UpdatedAt = time.Now()
	final_response := database.DB.Save(&user)
	if final_response.Error != nil {
		return false
	}
	return true
}

func ListUsers() ([]models.User, error) {
	var user []models.User
	result := database.DB.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
