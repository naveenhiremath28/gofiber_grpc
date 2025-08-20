package dbservice

import (
	"crud-grpc-gofiber/internal/database"
	"crud-grpc-gofiber/internal/database/dbmodels"
)

func GetUser(id int32) ([]dbmodels.User, error) {
	var user []dbmodels.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
